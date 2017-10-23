package webClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// PostDataToGoServer 推送业务服务器数据
func PostDataToGoServer(apiStr APIType, data []interface{}, isAsync bool) (responseObj *webServerObject.ResponseObject, err error) {
	if isAsync {
		go postDataToGoServer(apiStr, data)
		return nil, nil
	}

	return postDataToGoServer(apiStr, data)
}

// postDataToGoServer 推送业务服务器数据
func postDataToGoServer(apiStr APIType, data []interface{}) (responseObj *webServerObject.ResponseObject, err error) {
	requestObj := make(map[string]interface{})
	requestObj["Data"] = data

	// 记录错误日志
	defer func() {
		if err != nil {
			logTool.LogError(fmt.Sprintf("推送业务服务器数据失败，err:%s", err.Error()))
		} else {
			logTool.LogDebug(fmt.Sprintf("推送业务服务器数据成功，requestObj:%s，responseObj：%s", requestObj, responseObj))
		}
	}()

	requestByte, _ := json.Marshal(requestObj)

	//构造请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", configTool.GoServerWebAddress, apiStr), strings.NewReader(string(requestByte)))
	req.Header.Add("User-Agent", "fileWebClient")
	req.Header.Add("Referer", configTool.Referer)
	req.Close = true

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	responseObj = webServerObject.NewResponseObject()
	if err := json.Unmarshal(body, responseObj); err != nil {
		return nil, err
	}

	return responseObj, nil
}
