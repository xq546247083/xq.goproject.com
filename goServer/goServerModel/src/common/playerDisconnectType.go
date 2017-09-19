package common

//PlayerDisconnectType 玩家断开连接类型
type PlayerDisconnectType int

const (
	//FromForbid 来自于封号
	FromForbid PlayerDisconnectType = 1 + iota

	//FromSilent 来自于禁言
	FromSilent
)
