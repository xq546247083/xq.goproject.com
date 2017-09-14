package transferObject

// 转发的对象
type ForwardObject struct {
	// 转发的消息类型
	MessageType MessageType

	// 在SocketServer和SocketServerCenter之间传输的聊天消息对象
	ChatMessageObject *ChatMessageObject
}

func NewForwardObject(messageType MessageType, chatMessageObj *ChatMessageObject) *ForwardObject {
	return &ForwardObject{
		MessageType:       messageType,
		ChatMessageObject: chatMessageObj,
	}
}
