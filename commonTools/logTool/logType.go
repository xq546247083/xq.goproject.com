package logTool

//LogType 日志类型
type LogType int

//ToString 转换枚举为字符串
func (logType LogType) ToString() string {
	return data[logType]
}

const (
	//Info 消息
	Info LogType = iota

	//Debug 调试
	Debug

	//Warn 警告
	Warn

	//Error 错误
	Error

	//Fatal 致命错误
	Fatal
)

var data = [...]string{
	"Info",
	"Debug",
	"Warn",
	"Error",
	"Fatal",
}
