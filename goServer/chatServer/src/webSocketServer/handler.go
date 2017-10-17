package webSocketServer

import "xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"

// 处理器对象
type handler struct {
	// 消息类型
	messageType string

	// 方法定义
	handlerFunc func(*webSocketServerObject.RequestObject)
}

// 创建新的请求方法对象
// _messageType：消息类型
// _handlerFunc：方法定义
func newHandler(_messageType string, _handlerFunc func(*webSocketServerObject.RequestObject)) *handler {
	return &handler{
		messageType: _messageType,
		handlerFunc: _handlerFunc,
	}
}
