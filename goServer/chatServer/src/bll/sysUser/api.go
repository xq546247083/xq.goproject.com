package sysUser

import (
	"time"

	"xq.goproject.com/goServer/chatServer/src/rpcServer"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
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

// checkRequest 检测请求
func checkRequest(requestObject *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()

	//如果不是这几个方法，则要检测用户数据
	if requestObject.HTTPRequest.RequestURI != "/API/SysUser/Login" && requestObject.HTTPRequest.RequestURI != "/API/SysUser/Register" &&
		requestObject.HTTPRequest.RequestURI != "/API/SysUser/Identify" && requestObject.HTTPRequest.RequestURI != "/API/SysUser/Retrieve" {
		//根据用户名字判断过期时间
		userName, err := requestObject.GetStringVal("UserName")
		token, err2 := requestObject.GetStringVal("Token")
		if err != nil || err2 != nil {
			responseObj.SetResultStatus(webServerObject.DataError)
			return responseObj
		}

		if GetUserToken(userName) != token {
			responseObj.SetResultStatus(webServerObject.SignError)
			return responseObj
		}

		//如果过期，返回过期提示
		if CheckPwdExpiredTime(userName) {
			responseObj.SetResultStatus(webServerObject.LoginIsOverTime)
			return responseObj
		} else {
			//如果没过期，返回新的过期时间
			UpdatePwdExpiredTime(userName)
			sysUserObj := GetItemByUserNameOrEmail(userName)
			if sysUserObj != nil {
				responseObj.AttachData["PwdExpiredTime"] = sysUserObj.PwdExpiredTime.UnixNano() / 1e6
			}
		}
	}

	return responseObj
}
