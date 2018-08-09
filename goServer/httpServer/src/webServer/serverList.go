package webServer

import (
	"fmt"
	"net/http"
	"sync"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/logTool"
)

var (
	// 服务器列表
	serverList = make([]*http.Server, 0, 2)

	// 服务器锁
	serverMuext = new(sync.Mutex)
)

// 添加服务器
func addServer(serverItem *http.Server) {
	serverMuext.Lock()
	defer serverMuext.Unlock()

	serverList = append(serverList, serverItem)
}

// 移除服务器
func removeServer(serverItem *http.Server) {
	serverMuext.Lock()
	defer serverMuext.Unlock()

	for index, value := range serverList {
		if serverItem == value {
			serverList = append(serverList[:index], serverList[index+1:]...)
		}
	}
}

// 开启配置的服务器
func startAllServer() {
	if configTool.Protocol == "http" {
		go startServer(configTool.ListenPort)
	} else if configTool.Protocol == "https" {
		go startServerTLS(configTool.ListenPort)
	} else {
		go startServer(configTool.ListenPort)

		// 端口判定
		httpsPort := 443
		if configTool.ListenPort != 80 {
			httpsPort = configTool.ListenPort + 1
		}
		go startServerTLS(httpsPort)
	}
}

// 关闭所有开启的服务器
func closeAllServer() {
	serverMuext.Lock()
	defer serverMuext.Unlock()

	logTool.LogInfo(fmt.Sprintf("开始关闭服务器!"))
	for _, serverItem := range serverList {
		if err := serverItem.Close(); err != nil {
			logTool.LogError(fmt.Sprintf("关闭服务器出错！serverItem：%v.err:%s", serverItem, err))
		}
	}

	serverList = serverList[:0]
	logTool.LogInfo(fmt.Sprintf("关闭服务器成功!"))
}
