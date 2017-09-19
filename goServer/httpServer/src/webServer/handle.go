package webServer

import (
	"net/http"
	"path"
	"strings"
)

//Handle webserver服务处理
type handle struct{}

//服务监听
func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	upath := request.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		request.URL.Path = upath
	}

	serveFile(responseWriter, request, httpDir, path.Clean(upath), true)
}
