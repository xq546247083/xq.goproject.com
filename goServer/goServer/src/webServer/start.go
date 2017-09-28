package webServer

import (
	"xq.goproject.com/commonTools/logTool"
	"fmt"
	"net/http"
	"sync"
)

//StartServer 开启服务
func StartServer(wg *sync.WaitGroup, serverAddress string) {
	defer func() {
		wg.Done()
	}()

	//开启服务
	logTool.Log(logTool.Info, fmt.Sprintf("Web服务器监听：%s", serverAddress))
	fmt.Println( fmt.Sprintf("Web服务器监听：%s", serverAddress))

	if err := http.ListenAndServe(serverAddress, new(handle)); err != nil {
		logTool.LogObject(logTool.Error, err)
	}

}
