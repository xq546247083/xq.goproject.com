package main

import (
	"sync"

	"xq.goproject.com/commonTool/configTool"
	_ "xq.goproject.com/goServer/fileServer/src/bll"
	"xq.goproject.com/goServer/fileServer/src/webServer"
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
