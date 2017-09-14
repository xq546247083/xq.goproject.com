package playerBLL

import (
	"xq.goproject.com/goServer/src/rpcServer"
	"xq.goproject.com/goServer/src/webServer"
	"xq.goproject.com/goServerModel/src/player"
	"xq.goproject.com/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServerModel/src/webServerObject"
	"time"
)

// 注册需要给客户端访问的模块、方法
func init() {
	rpcServer.RegisterHandler("PlayerLogin", Login)
	webServer.RegisterHandler("PlayerLogin", WebLogin)
}

//Login 玩家登录
func Login(requestObj *rpcServerObject.RequestObject) *rpcServerObject.ResponseObject {
	responseObj := rpcServerObject.NewResponseObject()
	responseObj.SetResultStatus(rpcServerObject.Success)
	responseObj.Data = requestObj.Parameters[0]

	clientObj, ok := requestObj.Parameters[1].(*rpcServer.Client)
	if !ok {
		responseObj.Data = "转换client失败"
	}

	//给玩家注册客户端id
	player := player.NewEmptyPlayer("1dsad", clientObj.GetID())

	go func() {
		for{
			time.Sleep(10 * time.Second)
			clientObj := rpcServer.GetClient(player.ClientID)
			responseObj.SetResultStatus(rpcServerObject.Success)
			responseObj.Data = "推送消息"
	
			rpcServer.ResponseResult(clientObj, responseObj, rpcServer.ConHighPriority)
		}
	}()

	return responseObj
}

//WebLogin 玩家登录
func WebLogin(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)
	responseObj.Data = requestObj.Parameters[0]

	return responseObj
}

//Beat 心跳
func Beat(playerObj *player.Player) *rpcServerObject.ResponseObject {
	responseObj := rpcServerObject.NewResponseObject()

	return responseObj
}
