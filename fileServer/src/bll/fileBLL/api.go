package fileBLL

import (
	"fmt"
	"net/http"

	"xq.goproject.com/commonTool/httpRequestTool"
	"xq.goproject.com/fileServer/src/webServer"
	"xq.goproject.com/goServerModel/src/webServerObject"
)

var (
	uploadPath = "./upload/"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/UploadFile", uploadFile)
	webServer.RegisterHandler("/API/DownFile", downFile)
}

// 上传文件
func uploadFile(requestObj *http.Request) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)

	//从请求当中判断方法
	if requestObj.Method == "POST" {
		//获取文件内容 要这样获取
		file, head, err := requestObj.FormFile("file")
		if err != nil {
			responseObj.SetResultStatus(webServerObject.DataError)
			return responseObj
		}
		defer file.Close()

		//获取上传用户和时间
		userName := requestObj.FormValue("userName")
		picName := requestObj.FormValue("picName")
		uploadTime := requestObj.FormValue("uploadTime")
		if userName == "" || uploadTime == "" {
			responseObj.SetResultStatus(webServerObject.DataError)
			return responseObj
		}

		errTwo := saveFile(fmt.Sprintf("%s_%s_%s_%s", userName,picName, uploadTime, head.Filename), file)
		if errTwo != nil {
			responseObj.SetResultStatus(webServerObject.SaveFileFail)
			return responseObj
		}
	}

	return responseObj
}

//下载文件
func downFile(requestObj *http.Request) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)
	data, _ := httpRequestTool.GetRequsetByte(requestObj)
	responseObj.Data = string(data)

	return responseObj
}
