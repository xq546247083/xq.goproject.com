package main

import (
	"sync"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/goServer/httpServer/src/webServer"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(2)
}

func main() {
	//开启web服务
	go webServer.StartServer(&wg, configTool.WebListenAddress)

	go webServer.StartServer2(&wg, configTool.WebListenAddresss, configTool.Crt, configTool.Key)

	wg.Wait()
}
