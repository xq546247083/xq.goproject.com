package chat

import (
	"time"

	"xq.goproject.com/goServer/chatServer/src/rpcServer"
	"xq.goproject.com/goServer/chatServer/src/webSocketServer"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	rpcServer.RegisterHandler("RpcTest", rpcTest)

	webSocketServer.RegisterHandler("RpcTest", rpcTest1)
}

//rpcTest1 rpcTest方法
func rpcTest1(requestObj *webSocketServerObject.RequestObject) {
	responseObj := rpcServerObject.NewResponseObject()
	responseObj.SetResultStatus(rpcServerObject.Success)
	userName, err := requestObj.GetStringData(1)
	if err != nil {
		return
	}

	_ = userName
}

//rpcTest rpcTest方法
func rpcTest(requestObj *rpcServerObject.RequestObject) {
	clientObj, err := rpcServer.GetRequestClient(requestObj)
	if err != nil {
		return
	}

	responseObj := rpcServerObject.NewResponseObject()
	responseObj.SetResultStatus(rpcServerObject.Success)
	userName, err := requestObj.GetStringData(1)
	if err != nil {
		return
	}

	go func() {
		for {
			clientObj := rpcServer.GetClient(clientObj.GetID())
			responseObj.SetResultStatus(rpcServerObject.Success)
			responseObj.Data = userName

			rpcServer.ResponseResult(clientObj, responseObj, rpcServer.ConHighPriority)
			time.Sleep(10 * time.Second)
		}
	}()
}
