package rpcServer

import (
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
)

// 处理器对象
type handler struct {
	// 消息类型
	messageType string

	// 方法定义
	handlerFunc func(*rpcServerObject.RequestObject) *rpcServerObject.ResponseObject
}

// 创建新的请求方法对象
// _messageType：消息类型
// _handlerFunc：方法定义
func newHandler(_messageType string, _handlerFunc func(*rpcServerObject.RequestObject) *rpcServerObject.ResponseObject) *handler {
	return &handler{
		messageType: _messageType,
		handlerFunc: _handlerFunc,
	}
}
