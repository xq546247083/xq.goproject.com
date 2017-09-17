package fileBLL

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

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
		uploadTime := requestObj.FormValue("uploadTime")
		if userName == "" || uploadTime == "" {
			responseObj.SetResultStatus(webServerObject.DataError)
			return responseObj
		}

		errTwo := saveFile(fmt.Sprintf("%s_%s_%s", userName, uploadTime, head.Filename), file)
		if errTwo != nil {
			responseObj.SetResultStatus(webServerObject.SaveFileFail)
			return responseObj
		}
	}

	return responseObj
}

// downFile 下载文件
func downFile(requestObj *http.Request) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	responseObj.SetResultStatus(webServerObject.Success)
	data, _ := httpRequestTool.GetRequsetByte(requestObj)
	responseObj.Data = string(data)

	return responseObj
}

//保存文件
func saveFile(fileName string, fileReader multipart.File) error {
	//判断文件夹是否存在，如果不存在，则创建
	exists, err := isDirectoryExists(uploadPath)
	if err != nil {
		return err
	}

	if !exists {
		if err = os.Mkdir(uploadPath, os.ModePerm); err != nil {
			return err
		}
	}

	//文件全路径
	fileFullPath := uploadPath + fileName
	//判断文件夹是否存在，,如果存在，则追加
	fileExists, errT := isFileExists(fileFullPath)
	if errT != nil {
		return err
	}

	//如果不存在，则创建
	if !fileExists {
		//创建文件写入流
		fileWriter, err := os.Create(fileFullPath)
		if err != nil {
			return err
		}
		defer fileWriter.Close()

		//复制文件
		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}
	} else {
		errC := appendToFile(fileFullPath, fileReader)
		if errC != nil {
			return errC
		}
	}

	return nil
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, file multipart.File) error {
	// 以追加的模式，打开文件
	openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return err
	}

	// 查找文件末尾的偏移量
	n, err := openFile.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}

	//读取字节
	data, err := readMultiPartFile(file)
	if err != nil {
		return err
	}

	// 从末尾的偏移量开始写入内容
	len, err := openFile.WriteAt(data, n)
	if err != nil {
		return err
	}
	_ = len
	defer openFile.Close()

	return nil
}

func readMultiPartFile(file multipart.File) ([]byte, error) {
	len, errLen := file.Seek(0, os.SEEK_END)
	if errLen != nil {
		return nil, errLen
	}

	result := make([]byte, len, len)

	buf := make([]byte, 1024, 1024)
	for {
		//todo xiaoqiang  这里的字节数全是空
		nr, err := file.Read(buf)
		if nr > 0 {
			appendBuf := buf[:nr]
			buf = append(result, appendBuf...)
		}
		if err != nil {
			if err.Error() != "EOF" {
				return nil, err
			}
			break
		}
	}

	return result, nil
}

// 文件夹是否存在
func isDirectoryExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

// 文件是否存在
func isFileExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file.IsDir() == false, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}
