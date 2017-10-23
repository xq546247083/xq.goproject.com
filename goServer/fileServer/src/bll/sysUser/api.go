package sysUser

import (
	"encoding/json"

	"xq.goproject.com/goServer/fileServer/src/model"
	"xq.goproject.com/goServer/fileServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/Func/SysUser/UpdateUser", updateUser)
	webServer.RegisterHandler("/InnerFunc/SysUser/CheckHandler", checkRequest)
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
func checkRequest(requestObject *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()

	//从请求当中判断方法
	if requestObject.HTTPRequest.Method == "POST" {
		// 根据用户名字判断过期时间
		userName := requestObject.HTTPRequest.FormValue("UserName")
		token := requestObject.HTTPRequest.FormValue("Token")

		// 如果不正确，先从业务服务器获取数据，再校验
		if GetUserToken(userName) != token || CheckPwdExpiredTime(userName) {
			//如果获取数据失败，直接返回
			getUserResponse := getUserDataByGoServer(userName)
			if getUserResponse.Status != webServerObject.Success {
				return getUserResponse
			}

			//如果再校验失败，直接返回
			if GetUserToken(userName) != token || CheckPwdExpiredTime(userName) {
				responseObj.SetResultStatus(webServerObject.ClientDataError)
				return responseObj
			}
		}
	} else {
		responseObj.SetResultStatus(webServerObject.ClientDataError)
	}

	return responseObj
}
