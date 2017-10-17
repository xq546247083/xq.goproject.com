package sysUser

import (
	"encoding/json"

	"xq.goproject.com/goServer/chatServer/src/model"
	"xq.goproject.com/goServer/chatServer/src/webServer"
	"xq.goproject.com/goServer/chatServer/src/webSocketServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/SysUser/UpdateUser", updateUser)
	webSocketServer.RegisterCheckHandler(checkRequest)
}

// updateUser 更新用户数据
func updateUser(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	sysUserByte, err := requestObj.GetByteData(1)
	if err != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	// 转换用户数据
	sysUser := new(model.SysUser)
	err2 := json.Unmarshal(sysUserByte, sysUser)
	if err2 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	updateSysUser(sysUser)

	return responseObj
}

// checkRequest 检测请求
func checkRequest(requestObject *webSocketServerObject.RequestObject) bool {
	// 根据用户名字判断过期时间
	userName, err := requestObject.GetStringVal("UserName")
	token, err2 := requestObject.GetStringVal("Token")
	if err != nil || err2 != nil {
		return false
	}

	if GetUserToken(userName) != token {
		return false
	}

	// 如果过期，返回过期提示
	if CheckPwdExpiredTime(userName) {
		return false
	}

	return true
}
