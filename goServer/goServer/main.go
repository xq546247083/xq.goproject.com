package main

import (
	"sync"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/goServer/src/rpcServer"
	"xq.goproject.com/goServer/goServer/src/webServer"

	_ "github.com/go-sql-driver/mysql"
	_ "xq.goproject.com/goServer/goServer/src/bll"
)

var (
	wg sync.WaitGroup
)

func init() {
	logTool.LogInfo("开始启动服务器！！！")
	wg.Add(1)
}

func main() {
	//调用初始化和检测数据
	initTool.InitAndCheckData()

	//开启rpc服务
	go rpcServer.StartServer(&wg, configTool.RPCListenAddress)

	//开启web服务
	go webServer.StartServer(&wg, configTool.WebListenAddress)

	wg.Wait()
}
