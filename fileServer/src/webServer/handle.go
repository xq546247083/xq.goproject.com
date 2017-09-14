package webServer

import (
	"xq.goproject.com/goServerModel/src/webServerObject"
	"xq.goproject.com/commonTool/logTool"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"mime/multipart"
)

//Handle webserver服务处理
type handle struct{}

func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseObj := webServerObject.NewResponseObject() //返回的数据
	logInfo := ""                                      //最后写的日志

	//返回前，先返回数据，后写日志
	defer func() {
		data, err := json.Marshal(responseObj)
		if err != nil {
			logInfo = fmt.Sprintf("返回对象-反序列化出错，请求:%s错误信息为：%s", request, err)
			//返回对象反序列化失败，只能返回空数据
			data = []byte("")
		}

		responseWriter.Write(data)
		if logInfo != "" {
			logTool.Log(logTool.Error, logInfo)
		}else{
			logTool.Log(logTool.Debug, "web服务器返回数据："+string(data))
		}
	}()

	// 监控请求地址，必须为File
	if request.RequestURI != "/File" && request.RequestURI != "/File/" {
		return
	}

	//获取请求数据
	requestData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logInfo = fmt.Sprintf("读取数据出错，请求:%s,错误信息为：%s", request, err)
		responseObj.SetResultStatus(webServerObject.DataError)
		return
	}
	request.Body.Close()

	// 解析请求字符串
	if err := json.Unmarshal(requestData, &requestObj); err != nil {
		logInfo = fmt.Sprintf("请求对象-反序列化出错，请求:%s错误信息为：%s", request, err)
		responseObj.SetResultStatus(webServerObject.DataError)
		return
	}

	logTool.Log(logTool.Debug, "web服务器接受到请求："+string(requestData))
	responseObj = callFunction(requestObj.MethodName, &requestObj)
}
