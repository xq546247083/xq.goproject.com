package webClient

import (
	"xq.goproject.com/commonTools/configTool"
)

// ServerType 服务器类型
type ServerType string

var (
	// ChatWebServer 聊天服务器地址
	ChatWebServer = ServerType(configTool.ChatServerWebAddress)

	// FileWebServer 文件服务器地址
	FileWebServer = ServerType(configTool.FileServerAddress)
)
