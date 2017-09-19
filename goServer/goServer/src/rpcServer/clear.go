package rpcServer

import (
	"xq.goproject.com/goServer/goServerModel/src/common"
	"xq.goproject.com/goServer/goServer/src/goroutineMap"
	"xq.goproject.com/commonTool/logTool"
	"fmt"
	"time"
)

//clearClient 清理客户端
func clearClient() {
	goroutineMap.Operate("clearClient", common.AddOperate)
	defer goroutineMap.Operate("clearClient", common.ReduceOperate)
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
