package chat

import (
	"strings"
	"time"

	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/chatServer/src/bll/sysUser"
	"xq.goproject.com/goServer/chatServer/src/model"
	"xq.goproject.com/goServer/chatServer/src/rpcServer"
	"xq.goproject.com/goServer/chatServer/src/webSocketServer"
	"xq.goproject.com/goServer/goServerModel/src/consts"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	rpcServer.RegisterHandler("RpcTest", rpcTest)

	webSocketServer.RegisterHandler("Login", login)
	webSocketServer.RegisterHandler("Logout", logout)
	webSocketServer.RegisterHandler("BroadMessgae", broadMessgae)
	webSocketServer.RegisterHandler("BroadClients", broadClients)
	webSocketServer.RegisterHandler("SendMessgae", sendMessgae)
}

//broadMessgae 广播消息
func broadMessgae(requestObj *webSocketServerObject.RequestObject) {
	responseObj := webSocketServerObject.NewResponseObject(webSocketServerObject.World)

	userName, err := requestObj.GetStringData(1)
	message, err2 := requestObj.GetStringData(3)
	if err != nil && err2 != nil {
		return
	}

	if stringTool.IsEmpty(message) {
		return
	}

	//如果发送玩家不存在，则返回
	selfSysUser := sysUser.GetItemByUserNameOrEmail(userName)
	if selfSysUser == nil {
		return
	}

	//如果发送的词语包含了禁止的词语，则返回
	for _, word := range configWordForbidList {
		if strings.Contains(message, word.Word) {
			return
		}
	}

	//检测敏感词,并替换
	for _, word := range configWordSensitiveList {
		message = strings.Replace(message, word.Word, "***", -1)
	}

	//插入聊天消息到数据库
	historyWorld := model.NewHistoryWorld(message, "", selfSysUser.UserID, selfSysUser.UserName)
	savetHistoryWorld(historyWorld)

	responseObj.Data = historyWorld
	//广播消息
	webSocketServer.BroadMessage(responseObj)
}

//sendMessgae 发送消息
func sendMessgae(requestObj *webSocketServerObject.RequestObject) {
	responseObj := webSocketServerObject.NewResponseObject(webSocketServerObject.Private)

	userName, err := requestObj.GetStringData(1)
	talkToUserName, err2 := requestObj.GetStringData(2)
	message, err3 := requestObj.GetStringData(3)
	if err != nil && err2 != nil && err3 != nil {
		return
	}

	if stringTool.IsEmpty(message) {
		return
	}

	//如果发送玩家不存在，则返回
	selfSysUser := sysUser.GetItemByUserNameOrEmail(userName)
	if selfSysUser == nil {
		return
	}

	//如果接受玩家不存在，则返回
	talkToSysUser := sysUser.GetItemByUserNameOrEmail(talkToUserName)
	if talkToSysUser == nil {
		return
	}

	//如果发送的词语包含了禁止的词语，则返回
	for _, word := range configWordForbidList {
		if strings.Contains(message, word.Word) {
			return
		}
	}

	//检测敏感词,并替换
	for _, word := range configWordSensitiveList {
		message = strings.Replace(message, word.Word, "***", -1)
	}

	//插入聊天消息到数据库
	historyPrivate := model.NewHistoryPrivate(talkToSysUser.UserID, talkToSysUser.UserName, message, "", selfSysUser.UserID, selfSysUser.UserName)

	responseObj.Data = historyPrivate
	//给对和自己方发消息
	webSocketServer.SendMessage(selfSysUser.UserName, responseObj)
	if webSocketServer.IsOnline(talkToSysUser.UserName) {
		webSocketServer.SendMessage(talkToSysUser.UserName, responseObj)
		historyPrivate.IsSend = true
	}

	historyPrivateMap[historyPrivate.SysUserName] = append(historyPrivateMap[historyPrivate.SysUserName], historyPrivate)
	savetHistoryPrivate(historyPrivate)
}

//login 登录
func login(requestObj *webSocketServerObject.RequestObject) {
	userName, err := requestObj.GetStringData(1)
	if err != nil {
		return
	}

	//给登陆用户发送未接受的消息
	historyPrivates := getUnSendHistoryPrivateList(userName)
	for _, historyPrivate := range historyPrivates {
		responseObj := webSocketServerObject.NewResponseObject(webSocketServerObject.Private)
		responseObj.Data = historyPrivate
		webSocketServer.SendMessage(userName, responseObj)

		//保存发送记录
		historyPrivate.IsSend = true
		savetHistoryPrivate(historyPrivate)
	}

	webSocketServer.SetOnlineStatus(userName, true)
	//广播在线消息
	broadClients(requestObj)
}

//logout 退出
func logout(requestObj *webSocketServerObject.RequestObject) {
	userName, err := requestObj.GetStringData(1)
	if err != nil {
		return
	}

	webSocketServer.SetOnlineStatus(userName, false)
	//广播在线消息
	broadClients(requestObj)
}

//broadClients 广播所有客户端
func broadClients(requestObj *webSocketServerObject.RequestObject) {
	responseObj := webSocketServerObject.NewResponseObject(webSocketServerObject.BroadClients)

	sysUsers := sysUser.GetAllSysUser()
	returnMap := make([]map[string]interface{}, 0, len(sysUsers))

	//循环用户，返回用户的数据
	for _, sysUser := range sysUsers {
		userStrMap := make(map[string]interface{})
		userStrMap[consts.UserName] = sysUser.UserName
		userStrMap[consts.FullName] = sysUser.FullName
		userStrMap[consts.Online] = webSocketServer.IsOnline(sysUser.UserName)

		returnMap = append(returnMap, userStrMap)
	}

	responseObj.Data = returnMap
	//广播消息
	webSocketServer.BroadMessage(responseObj)
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
