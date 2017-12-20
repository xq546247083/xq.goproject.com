package main

// 导入外部包
import (
	_ "github.com/go-sql-driver/mysql"
	_ "xq.goproject.com/goServer/goServer/src/bll"
)

// 导入包
import (
	"sync"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/logTool"
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
