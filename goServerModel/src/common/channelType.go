package common

//ChannelType 聊天频道类型定义
type ChannelType int

//ToString 获取枚举的字符串
func (ct ChannelType) ToString() string {
	return data[ct+1]
}

const (
	//System 系统频道
	System ChannelType = iota

	//World 世界频道
	World

	//Union 公会频道
	Union

	//Private 私聊频道
	Private

	//Avatar Avatar频道，用于游戏服务器通知客户端
	Avatar

	//CrossServer 跨服频道，用于跨服聊天
	CrossServer
)

//通道描述
var data = [...]string{
	"系统频道",
	"世界频道",
	"公会频道",
	"私聊频道",
	"Avatar频道",
	"跨服频道",
}
