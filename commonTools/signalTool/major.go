package signalMgr

import (
	"os"
	"os/signal"
	"syscall"

	"xq.goproject.com/commonTools/logTool"
)

// Start ...启动信号管理器
func Start() {
	go func() {
		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

		for {
			// 准备接收信息
			sig := <-sigs

			if sig == syscall.SIGHUP {
				logTool.LogInfo("收到重启的信号，准备重新加载配置")

				// 重新加载

				logTool.LogInfo("收到重启的信号，重新加载配置完成")
			} else {
				logTool.LogInfo("收到退出程序的信号，开始退出……")

				// 调用退出的方法

				logTool.LogInfo("收到退出程序的信号，退出完成……")

				// 一旦收到信号，则表明管理员希望退出程序，则先保存信息，然后退出
				os.Exit(0)
			}
		}
	}()
}
