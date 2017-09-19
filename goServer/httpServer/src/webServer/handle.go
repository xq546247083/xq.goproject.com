package webServer

import (
	"net/http"
	"strings"

	"xq.goproject.com/commonTool/configTool"
)

var (
	webMainPath = configTool.WebMainPath
)

//Handle webserver服务处理
type handle struct{}

//服务监听
func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	requestType := path[strings.LastIndex(path, "."):]
	switch requestType {
	case ".css":
		responseWriter.Header().Set("content-type", "text/css")
	case ".js":
		responseWriter.Header().Set("content-type", "text/javascript")
	default:
	}

}
