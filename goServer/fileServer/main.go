package main

import (
	"sync"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/goServer/fileServer/src/webServer"

	_ "xq.goproject.com/goServer/fileServer/src/bll"
)

var (
	wg sync.WaitGroup
)

func init() {
	logTool.LogInfo("开始启动服务器！！！")

	wg.Add(1)
}

func main() {
	//设置邮箱
	emailTool.SetSenderInfo(configTool.EmailHost, configTool.EmailPort, configTool.EmailName, configTool.EmailAddress, configTool.EmailPass)

	//调用初始化和检测数据
	initTool.InitAndCheckData()

	//开启web服务
	go webServer.StartServer(&wg, configTool.WebListenAddress)

	wg.Wait()
}
