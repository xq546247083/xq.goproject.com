package transferObject

type TransferType string

const (
	// 登陆
	Login TransferType = "Login"

	// 转发信息
	Forward TransferType = "Forward"

	// 更新客户端和玩家数量
	UpdateClientAndPlayerCount TransferType = "UpdateClientAndPlayerCount"

	// 玩家登陆中心服务器
	PlayerLogin TransferType = "PlayerLogin"

	// 玩家登出中心服务器
	PlayerLogout TransferType = "PlayerLogout"
)
