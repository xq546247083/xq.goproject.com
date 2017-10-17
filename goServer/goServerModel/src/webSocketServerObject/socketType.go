package webSocketServerObject

// SocketType 返回的socket类型
type SocketType string

var (
	// BroadClients 广播客户端
	BroadClients SocketType = "BroadClients"

	// World 世界聊天
	World SocketType = "World"

	// Private 私密聊天
	Private SocketType = "Private"
)
