package transferObject

// 转发的消息类型
type MessageType string

const (
	// 聊天消息
	ChatMessage MessageType = "ChatMessage"

	// 推送消息
	PushMessage MessageType = "PushMessage"

	// 封号
	Forbid MessageType = "Forbid"

	// 禁言
	Silent MessageType = "Silent"

	// 重新加载配置
	Reload MessageType = "Reload"
)
