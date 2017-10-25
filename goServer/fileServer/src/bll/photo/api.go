package photo

import (
	"fmt"
	"time"

	"xq.goproject.com/goServer/fileServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/consts"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/APIFromFile/UploadPhoto", uploadPhoto)

	webServer.RegisterHandler("/API/DownPhoto", downPhoto)
	webServer.RegisterHandler("/API/Photo/GetUserPhotos", getUserPhotos)
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

	fileName := fmt.Sprintf("%s_%s_%s_%s", userName, picName, uploadTime, head.Filename)
	errTwo := saveFile(fileName, file)
	if errTwo != nil {
		responseObj.SetResultStatus(webServerObject.SaveFileFail)
		return responseObj
	}

	addPhoto(userName, ablum, fileName, time.Now())

	//返会结果
	clientInfo := make(map[string]interface{})

	clientInfo[consts.DirName] = uploadPath
	clientInfo[consts.FileName] = fileName

	responseObj.Data = clientInfo
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
