package rpcServer

import (
	"fmt"
	"time"

	"xq.goproject.com/commonTool/goroutineTool"
	"xq.goproject.com/commonTool/logTool"
	"xq.goproject.com/goServer/goServerModel/src/common"
)

//clearClient 清理客户端
func clearClient() {
	goroutineTool.Operate("clearClient", common.AddOperate)
	defer goroutineTool.Operate("clearClient", common.ReduceOperate)

	for {
		//由于300s为离开活跃的时间，所以300s清理一次
		time.Sleep(300 * time.Second)

		clientList := getObsoleteClient()

		for _, item := range clientList {
			item.Quit()
		}

		logTool.Log(logTool.Info, fmt.Sprintf("当前在线玩家为：%d,清理的不活跃玩家为：%d", getClientCount(), len(clientList)))
	}
}
