package photo

import (
	"fmt"

	"xq.goproject.com/goServer/fileServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/APIFromFile/UploadPhoto", uploadPhoto)
	webServer.RegisterHandler("/APIFromFile/DownPhoto", downPhoto)
	webServer.RegisterHandler("/API/GetUserPhotos", getUserPhotos)
}

// 上传文件
func uploadPhoto(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)

	//获取文件内容 要这样获取
	file, head, err := requestObj.HTTPRequest.FormFile("file")
	if err != nil {
		responseObj.SetResultStatus(webServerObject.DataError)
		return responseObj
	}
	defer file.Close()

	//获取上传用户和时间
	userName := requestObj.HTTPRequest.FormValue("UserName")
	picName := requestObj.HTTPRequest.FormValue("PicName")
	uploadTime := requestObj.HTTPRequest.FormValue("UploadTime")
	if userName == "" || uploadTime == "" {
		responseObj.SetResultStatus(webServerObject.DataError)
		return responseObj
	}

	errTwo := saveFile(fmt.Sprintf("%s_%s_%s_%s", userName, picName, uploadTime, head.Filename), file)
	if errTwo != nil {
		responseObj.SetResultStatus(webServerObject.SaveFileFail)
		return responseObj
	}
	return responseObj
}

// 获取用户照片
func getUserPhotos(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userName, err := requestObj.GetStringData(1)
	if err != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	responseObj.Data = assembleToClient(userName)
	return responseObj
}

//下载文件
func downPhoto(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()

	return responseObj
}