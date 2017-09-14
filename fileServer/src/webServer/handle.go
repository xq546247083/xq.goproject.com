package webServer

import (
	"xq.goproject.com/goServerModel/src/webServerObject"
	"xq.goproject.com/commonTool/logTool"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io"
)

//Handle webserver服务处理
type handle struct{}

func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseObj := webServerObject.NewResponseObject() //返回的数据
	logInfo := ""                                      //最后写的日志

	//返回前，先返回数据，后写日志
	defer func() {
		data, err := json.Marshal(responseObj)
		if err != nil {
			logInfo = fmt.Sprintf("返回对象-反序列化出错，请求:%s错误信息为：%s", request, err)
			//返回对象反序列化失败，只能返回空数据
			data = []byte("")
		}

		responseWriter.Header().Add("Access-Control-Allow-Origin","*")
		responseWriter.Write(data)
		if logInfo != "" {
			logTool.Log(logTool.Error, logInfo)
		}else{
			logTool.Log(logTool.Debug, "web服务器返回数据："+string(data))
		}
	}()

	// 监控请求地址，必须为File
	if request.RequestURI != "/File" && request.RequestURI != "/File/" {
		return
	}

	if "POST" == request.Method {
		uploadFile, _, err := request.FormFile("userfile")
		if err != nil {
			logInfo = fmt.Sprintf("保存文件失败", request, err)
			responseObj.SetResultStatus(webServerObject.DataError)
			return
		}
		defer uploadFile.Close()
		saveFile,err:=os.Create("filenametosaveas")
		defer saveFile.Close()
		io.Copy(saveFile,uploadFile)
		//fmt.Println(responseWriter, "上传文件的大小为: %d", uploadFile.(Sizer).Size())
		return
    }

	logTool.Log(logTool.Debug, "web服务器接受到请求：")
	//responseObj = callFunction(requestObj.MethodName, &requestObj)
}
