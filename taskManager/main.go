package main

// 导入包
import (
	"sync"

	"xq.goproject.com/taskManager/src/task"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
	"xq.goproject.com/commonTools/logTool"
)

var (
	wg sync.WaitGroup
)

func init() {
	logTool.LogInfo("开始启动！！！")
	wg.Add(1)
}

func main() {
	// 设置邮箱信息
	emailTool.SetSenderInfo(configTool.EmailHost, configTool.EmailPort, configTool.EmailName, configTool.EmailAddress, configTool.EmailPass)
	
	go task.Moniter()
	wg.Wait()
}
