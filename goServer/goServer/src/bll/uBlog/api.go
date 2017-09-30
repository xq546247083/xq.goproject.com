package uBlog

import (
	"xq.goproject.com/goServer/goServer/src/bll/sysUser"
	"xq.goproject.com/goServer/goServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/UBlog/GetBlogList", getBlogList)
}

//getBlogList 获取博客信息
func getBlogList(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userName, err := requestObj.GetStringVal("UserName")
	blogType, err2 := requestObj.GetInt32Val("BlogType")
	status, err3 := requestObj.GetInt32Val("Status")
	tagInfo, err4 := requestObj.GetStringVal("TagInfo")
	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//获取用户
	sysUser := sysUser.GetItemByUserNameOrEmail(userName)
	if sysUser == nil {
		responseObj.SetResultStatus(webServerObject.UserIsNotExist)
		return responseObj
	}

	//返回用户菜单信息
	responseObj.Data = assembleToClient(sysUser, blogType, status, tagInfo)

	return responseObj
}
