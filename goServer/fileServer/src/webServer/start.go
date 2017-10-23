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

	//服务获取文件
	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	//处理其他消息
	http.Handle("/", new(handle))

	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		logTool.LogObject(logTool.Error, err)
	}
}
