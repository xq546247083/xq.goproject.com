package webServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

//Handle webserver服务处理
type handle struct{}

//服务监听
func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseObj := webServerObject.NewResponseObject()
	logInfo := ""

	// 应对NLB的监控
	if request.RequestURI == "/" || request.RequestURI == "/favicon.ico" {
		return
	}

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
		} else {
			logTool.Log(logTool.Debug, fmt.Sprintf("web服务器接受请求：%s返回数据：%s", request, string(data)))
		}
	}()

	// 根据路径选择不同的处理方法
	handlerObj, exists := getHandler(request.RequestURI)
	if !exists {
		responseObj.SetResultStatus(webServerObject.DataError)
		return
	}

	// 调用方法
	responseObj = handlerObj.handlerFunc(request)
}
