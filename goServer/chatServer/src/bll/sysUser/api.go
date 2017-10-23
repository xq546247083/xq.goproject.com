package sysUser

import (
	"encoding/json"
	"fmt"

	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/chatServer/src/model"
	"xq.goproject.com/goServer/chatServer/src/webClient"
	"xq.goproject.com/goServer/chatServer/src/webServer"
	"xq.goproject.com/goServer/chatServer/src/webSocketServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/Func/SysUser/UpdateUser", updateUser)
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

	//如果是登录，先去业务服务器获取用户最新数据
	if requestObject.MethodName == "Login" {
		// 获取用户数据
		responseObj, err3 := webClient.PostDataToGoServer(webClient.GetUser, []interface{}{userName}, false)
		if err3 != nil {
			logTool.LogError(fmt.Sprintf("登录聊天服务器，拉取用户数据失败，err:%s", err3))
			return false
		}

		if responseObj.Data == nil {
			responseObj.SetResultStatus(webServerObject.ClientDataError)
			return false
		}

		//反序列化字典为byte
		dataByte, err4 := json.Marshal(responseObj.Data)
		if err4 != nil {
			logTool.LogError(fmt.Sprintf("登录聊天服务器，拉取用户数据，序列化失败，err:%s", err4))
			return false
		}

		getUser := &model.SysUser{}
		//再序列化为对象
		err5 := json.Unmarshal(dataByte, getUser)
		if err5 != nil {
			logTool.LogError(fmt.Sprintf("登录聊天服务器，拉取用户数据，反序列化失败，err:%s", err5))
			return false
		}

		//更新玩家数据
		sysUserMap[getUser.UserName] = getUser
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
