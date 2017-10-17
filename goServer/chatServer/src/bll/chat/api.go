package chat

import (
	"strings"
	"time"

	"xq.goproject.com/goServer/chatServer/src/bll/sysUser"
	"xq.goproject.com/goServer/chatServer/src/model"
	"xq.goproject.com/goServer/chatServer/src/rpcServer"
	"xq.goproject.com/goServer/chatServer/src/webSocketServer"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webSocketServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	rpcServer.RegisterHandler("RpcTest", rpcTest)

	webSocketServer.RegisterHandler("SendMessgaeInWorld", sendMessgaeInWorld)
	webSocketServer.RegisterHandler("BroadClients", broadClients)
}

//sendMessgaeInWorld 广播消息
func sendMessgaeInWorld(requestObj *webSocketServerObject.RequestObject) {
	responseObj := webSocketServerObject.NewResponseObject(webSocketServerObject.World)

	userName, err := requestObj.GetStringData(1)
	message, err2 := requestObj.GetStringData(2)
	if err != nil && err2 != nil {
		return
	}

	//如果发送玩家不存在，则返回
	sysUser := sysUser.GetItemByUserNameOrEmail(userName)
	if sysUser == nil {
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
	historyWorld := model.NewHistoryWorld(message, "", sysUser.UserID, sysUser.UserName)
	insertHistoryWorld(historyWorld)

	responseObj.Data = historyWorld
	//广播消息
	webSocketServer.BroadMessage(responseObj)
}

//broadClients 广播所有客户端
func broadClients(requestObj *webSocketServerObject.RequestObject) {
	responseObj := webSocketServerObject.NewResponseObject(webSocketServerObject.BroadClients)

	clientUserNames := webSocketServer.GetAllClientUserName()
	returnMap := make([]map[string]string, 0, len(clientUserNames))

	//循环客户端用户名，添加用户的信息，返回
	for _, clientUserName := range clientUserNames {
		sysUser := sysUser.GetItemByUserNameOrEmail(clientUserName)
		if sysUser != nil {
			userStrMap := make(map[string]string)
			userStrMap["UserName"] = sysUser.UserName
			userStrMap["FullName"] = sysUser.FullName

			returnMap = append(returnMap, userStrMap)
		}
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
