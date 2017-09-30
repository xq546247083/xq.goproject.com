package sysUser

import (
	"encoding/json"
	"time"

	"xq.goproject.com/goServer/goServer/src/rpcServer"
	"xq.goproject.com/goServer/goServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	rpcServer.RegisterHandler("PlayerLogin", Login)
	webServer.RegisterHandler("/API/PlayerLogin", WebLogin)
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

	go func() {
		for {
			time.Sleep(10 * time.Second)
			clientObj := rpcServer.GetClient(clientObj.GetID())
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
	//data, _ := httpRequestTool.GetRequsetByte(requestObj)
	str, _ := json.Marshal(sysUserMap)

	responseObj.Data = string(str)

	return responseObj
}
