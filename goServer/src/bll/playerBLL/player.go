package playerBLL

import (
	"xq.goproject.com/goServer/src/rpcServer"
	"xq.goproject.com/goServerModel/src/player"
	"xq.goproject.com/goServerModel/src/rpcServerObject"
)

//GetPlayerByID 根据玩家Id获得玩家对象
// clientObj：客户端对象
// playerId：玩家Id
// 返回值：
// 玩家对象
// 返回值枚举值
func GetPlayerByID(clientobj *rpcServer.Client, playerID string) (playerobj *player.Player, rs rpcServerObject.ResultStatus) {
	rs = rpcServerObject.APIDataError

	//创建一个player对象用于返回
	playerobj = player.NewEmptyPlayer(playerID, clientobj.GetID())

	return
}
