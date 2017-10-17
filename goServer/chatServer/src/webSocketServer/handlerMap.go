package webSocketServer

import (
	"fmt"

	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

var (
	// 所有对外提供的方法列表
	handlerMap = make(map[string]*handler)

	// 检测handler
	checkHandler func(*webSocketServerObject.RequestObject) bool
)

func init() {
	initTool.RegisterInitFunc(logRegisterHandlerInfo, initTool.I_Global)
}

// 记录注册的数据信息
func logRegisterHandlerInfo() error {
	handlerInfo := make([]string, 0, len(handlerMap))
	if len(handlerMap) > 0 {
		handlerInfo = append(handlerInfo, "webSocketServer当前已注册接口：")
	} else {
		handlerInfo = append(handlerInfo, "webSocketServer暂无注册接口")
	}

	//获取注册的接口名
	for name := range handlerMap {
		handlerInfo = append(handlerInfo, fmt.Sprintf("%s", name))
	}

	logTool.LogInfo(handlerInfo...)
	return nil
}

//RegisterHandler 注册方法
// methodName:调用方法全名
// handlerFunc：方法定义
func RegisterHandler(methodName string, handlerFunc func(*webSocketServerObject.RequestObject)) {
	if _, exists := handlerMap[methodName]; exists {
		panic(fmt.Sprintf("%s已经存在，请重新取名", methodName))
	}

	handlerMap[methodName] = newHandler(methodName, handlerFunc)
}

//调用方法
func callFunction(requestObj *webSocketServerObject.RequestObject) {
	handlerObj, exists := getHandler(requestObj.MethodName)
	// 查找方法
	if !exists {
		logTool.Log(logTool.Error, fmt.Sprintf("方法名:%s未定义", requestObj.MethodName))
		return
	}

	handlerObj.handlerFunc(requestObj)
}

//获取处理
func getHandler(methodName string) (*handler, bool) {
	if handlerObj, exists := handlerMap[methodName]; exists {
		return handlerObj, exists
	}

	return nil, false
}

//RegisterCheckHandler 注册方法
// methodName:调用方法全名
// handlerFunc：方法定义
func RegisterCheckHandler(handlerFunc func(*webSocketServerObject.RequestObject) bool) {
	if checkHandler != nil {
		panic("检测方法【checkHandler】已经存在!")
	}

	checkHandler = handlerFunc
}
