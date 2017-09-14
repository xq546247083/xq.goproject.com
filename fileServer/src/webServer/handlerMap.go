package webServer

import (
	"xq.goproject.com/goServerModel/src/webServerObject"
	"xq.goproject.com/commonTool/logTool"
	"fmt"
)

var (
	// 所有对外提供的方法列表
	handlerMap = make(map[string]*handler)
)

//RegisterHandler 注册方法
// methodName:调用方法全名
// handlerFunc：方法定义
func RegisterHandler(methodName string, handlerFunc func(*webServerObject.RequestObject) *webServerObject.ResponseObject) {
	if _, exists := handlerMap[methodName]; exists {
		panic(fmt.Sprintf("%s已经存在，请重新取名", methodName))
	}

	handlerMap[methodName] = newHandler(methodName, handlerFunc)
	logTool.Log(logTool.Info, fmt.Sprintf("WebSerber 注册方法:%s，当前共有%d个注册", methodName, len(handlerMap)))
}

//调用方法
func callFunction(requestURI string, requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()

	handlerObj, exists := getHandler(requestURI)
	// 查找方法
	if !exists {
		logTool.Log(logTool.Error, fmt.Sprintf("方法名:%s未定义", requestURI))

		return responseObj.SetResultStatus(webServerObject.NoTargetMethod)
	}

	return handlerObj.handlerFunc(requestObj)
}

//获取处理
func getHandler(methodName string) (*handler, bool) {
	if handlerObj, exists := handlerMap[methodName]; exists {
		return handlerObj, exists
	}

	return nil, false
}
