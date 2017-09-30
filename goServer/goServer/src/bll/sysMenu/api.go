package sysMenu

import (
	"encoding/json"

	"xq.goproject.com/commonTools/httpRequestTool"
	"xq.goproject.com/goServer/goServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/SysMenu/GetInfo", WebLogin)
}

//WebLogin 玩家登录
func WebLogin(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()

	data, _ := httpRequestTool.GetRequsetByte(requestObj)
	str, _ := json.Marshal(sysUserMap)

	responseObj.Data = string(str)

	return responseObj
}
