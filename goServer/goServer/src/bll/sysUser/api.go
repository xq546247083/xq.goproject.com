package sysUser

import (
	"time"

	"xq.goproject.com/commonTools/EncrpytTool"
	"xq.goproject.com/goServer/goServer/src/rpcServer"
	"xq.goproject.com/goServer/goServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/SysUser/Login", login)
	rpcServer.RegisterHandler("RpcTest", rpcTest)
}

//rpcTest rpcTest方法
func rpcTest(requestObj *rpcServerObject.RequestObject) *rpcServerObject.ResponseObject {
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

//login 获取菜单信息
func login(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userName, err := requestObj.GetStringData(1)
	userPwd, err2 := requestObj.GetStringData(1)
	if err != nil || err2 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//获取用户
	sysUser := GetItemByUserNameOrEmail(userName)
	if sysUser == nil {
		responseObj.SetResultStatus(webServerObject.UserIsNotExist)
		return responseObj
	}

	if userPwd == "6fda14112d9151ebefc40a96c9b85be3" {
		responseObj.SetResultStatus(webServerObject.PlsEnterPassword)
		return responseObj
	}

	if sysUser.Password != EncrpytTool.Encrypt(userPwd) {
		responseObj.SetResultStatus(webServerObject.PlsEnterPassword)
		return responseObj
	}

	//返回用户菜单信息
	responseObj.Data = assembleToClient(sysUser)

	return responseObj
}
