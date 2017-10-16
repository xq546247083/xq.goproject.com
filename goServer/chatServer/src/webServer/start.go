package webServer

import (
	"fmt"
	"net/http"
	"sync"

	"xq.goproject.com/commonTools/logTool"
)

//StartServer 开启服务
func StartServer(wg *sync.WaitGroup, serverAddress string) {
	defer func() {
		wg.Done()
	}()

	//开启服务
	logTool.Log(logTool.Info, fmt.Sprintf("Web服务器监听：%s", serverAddress))
	fmt.Println(fmt.Sprintf("Web服务器监听：%s", serverAddress))

	if err := http.ListenAndServe(serverAddress, new(handle)); err != nil {
		logTool.LogObject(logTool.Error, err)
	}
}
