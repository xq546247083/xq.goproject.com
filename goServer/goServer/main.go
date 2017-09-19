package main

import (
	_ "xq.goproject.com/goServer/goServer/src/bll"
	"xq.goproject.com/goServer/goServer/src/rpcServer"
	"xq.goproject.com/goServer/goServer/src/webServer"
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
	//开启rpc服务
	go rpcServer.StartServer(&wg,  configTool.RPCListenAddress)

	//开启web服务
	go webServer.StartServer(&wg, configTool.WebListenAddress)

	wg.Wait()
}
