package main

import (
	"sync"

	"xq.goproject.com/goServer/httpServer/src/updatePackage"
	"xq.goproject.com/goServer/httpServer/src/webServer"
)

var (
	// wg
	wg sync.WaitGroup

	// 监控服务通道
	monitorServerChan = make(chan bool, 0)
)

func init() {
	wg.Add(1)
}

func main() {
	go webServer.Monitor(monitorServerChan)
	go updatePackage.Monitor(monitorServerChan)

	// 开启服务
	monitorServerChan <- true
	wg.Wait()
}
