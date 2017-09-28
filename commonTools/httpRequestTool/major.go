package httpRequestTool

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"xq.goproject.com/commonTools/logTool"
)

//GetRequsetByte 获取请求的字段
func GetRequsetByte(requestObj *http.Request) ([]byte, error) {
	defer func() {
		requestObj.Body.Close()
	}()

	data, err := ioutil.ReadAll(requestObj.Body)
	if err != nil {
		logTool.Log(logTool.Error, fmt.Sprintf("url:%s,读取数据出错，错误信息为：%s", requestObj.RequestURI, err))
		return nil, err
	}

	return data, nil
}
