package sysUser

import (
	"time"

	"xq.goproject.com/goServer/chatServer/src/rpcServer"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	rpcServer.RegisterHandler("/API/SysUser/UpdateUser", updateUser)
	rpcServer.RegisterCheckHandler(checkRequest)
}

//updateUser 更新用户信息
func updateUser(requestObj *rpcServerObject.RequestObject) {
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

// checkRequest 检测请求
func checkRequest(requestObject *rpcServerObject.RequestObject) bool {
	//根据用户名字判断过期时间
	userName, err := requestObject.GetStringVal("UserName")
	token, err2 := requestObject.GetStringVal("Token")
	if err != nil || err2 != nil {
		return false
	}

	if GetUserToken(userName) != token {
		return false
	}

	//如果过期，返回过期提示
	if CheckPwdExpiredTime(userName) {
		return false
	}

	return true
}
