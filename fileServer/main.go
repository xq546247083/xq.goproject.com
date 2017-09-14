package main

import (
	_ "xq.goproject.com/fileServer/src/bll"
	"xq.goproject.com/fileServer/src/webServer"
	"xq.goproject.com/commonTool/configTool"
	"sync"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(1)
}

func main() {
	//开启web服务
	go webServer.StartServer(&wg, configTool.WebListenAddress)

	wg.Wait()
}
