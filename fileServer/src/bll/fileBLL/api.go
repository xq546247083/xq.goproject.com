package playerBLL

import (
	"xq.goproject.com/fileServer/src/webServer"
	"xq.goproject.com/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("UploadFile", uploadFile)
	webServer.RegisterHandler("DownFile", downFile)
}

// 上传文件
func uploadFile(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)
	responseObj.Data = requestObj.Parameters[0]

	return responseObj
}

// downFile 下载文件
func downFile(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)
	responseObj.Data = requestObj.Parameters[0]

	return responseObj
}
