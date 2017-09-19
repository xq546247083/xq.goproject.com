package webServer

import (
	"net/http"
)

//Handle webserver服务处理
type handle struct{}

//服务监听
func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

}
