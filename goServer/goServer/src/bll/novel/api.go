package novel

import (
	"xq.goproject.com/goServer/goServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/Novel/GetNovelList", getNovelList)
	webServer.RegisterHandler("/API/Novel/GetChapterList", getChapterList)
	webServer.RegisterHandler("/API/Novel/GetNovelInfo", getNovelInfo)
}

//getNovelList 获取小说列表
func getNovelList(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.Data = assembleNovelListToClient()

	return responseObj
}

//getChapterList 获取章节信息
func getChapterList(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	novelName, err := requestObj.GetStringData(1)
	num, err2 := requestObj.GetInt32Data(2)
	if err != nil || err2 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	responseObj.Data = assembleChapterListToClient(novelName, num)

	return responseObj
}

//getNovelInfo 获取内容
func getNovelInfo(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	novelName, err := requestObj.GetStringData(1)
	title, err2 := requestObj.GetStringData(2)
	flag, err2 := requestObj.GetInt32Data(3)
	if err != nil || err2 != nil || (flag != -1 && flag != 0 && flag != 1) {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	responseObj.Data = GetNovelInfos(novelName, title, flag)

	return responseObj
}
