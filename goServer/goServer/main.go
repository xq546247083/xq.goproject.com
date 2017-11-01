package main

//导入额外包
import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/logTool"
	_ "xq.goproject.com/goServer/goServer/src/bll"
	"xq.goproject.com/goServer/goServer/src/spider"
	"xq.goproject.com/goServer/goServer/src/webServer"
)

var (
	wg sync.WaitGroup
)

func init() {
	logTool.LogInfo("开始启动服务器！！！")
	wg.Add(1)
}

func main() {
	// 设置邮箱信息
	emailTool.SetSenderInfo(configTool.EmailHost, configTool.EmailPort, configTool.EmailName, configTool.EmailAddress, configTool.EmailPass)

	// 调用初始化和检测数据
	initTool.InitAndCheckData()

	// 开启web服务
	go webServer.StartServer(&wg, configTool.WebListenAddress)

	// 开始抓数据
	go spider.Start()

	wg.Wait()
}
