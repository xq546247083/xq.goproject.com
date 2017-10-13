package webServer

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/goServer/goServerModel/src/consts"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

//Handle webserver服务处理
type handle struct{}

//服务监听
func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	//如果是获取服务器配置
	if request.RequestURI == "/GetConfig" {
		getConfig(responseWriter, request)
		return
	}

	upath := request.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		request.URL.Path = upath
	}

	serveFile(responseWriter, request, httpDir, path.Clean(upath), true)
}

//getConfig 获取配置
func getConfig(responseWriter http.ResponseWriter, requestObj *http.Request) {
	responseObj := webServerObject.NewResponseObject()
	defer func() {
		data, err := json.Marshal(responseObj)
		if err != nil {
			//返回对象反序列化失败，返回空数据
			data = []byte("")
		}

		//返回数据
		responseWriter.Write(data)
	}()

	//返回用户信息
	clientInfo := make(map[string]interface{})
	clientInfo[consts.FileServerAddress] = configTool.FileServerAddress
	clientInfo[consts.GoServerAddress] = configTool.GoServerAddress

	responseObj.Data = clientInfo
}
