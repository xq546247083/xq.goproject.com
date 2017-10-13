package configTool

var (
	// LogPath 日志路径
	LogPath = "Log"

	// RPCListenAddress 监听地址
	RPCListenAddress string

	// WebListenAddress 监听地址
	WebListenAddress string

	// DBConnection 数据库地址
	DBConnection string

	// IsDebug 是否测试模式
	IsDebug = false

	// LogInfoFlag 是否记录消息
	LogInfoFlag = false

	// LogDebugFlag 是否记录Debug消息
	LogDebugFlag = false

	// LogWarnFlag 是否记录警告消息
	LogWarnFlag = false

	// LogErrorFlag 是否记录错误消息
	LogErrorFlag = false

	// LogFatalFlag 是否记录致命错误消息消息
	LogFatalFlag = false

	// WebMainPath 网站路径
	WebMainPath = "WebMain"

	// WebSiteReferer 网站标识
	WebSiteReferer = "xiaohe.info"

	// IndexPage 网站首页
	IndexPage = "/index.html"

	// Error404Page 404错误页面
	Error404Page = "/404.html"

	// Error500Page 500错误额亚明
	Error500Page = "/500.html"

	// PwdExpiredTime 密码过期时间
	PwdExpiredTime = 1

	// EmailHost 邮箱主机
	EmailHost = "smtp.qq.com"

	// EmailPort 邮箱端口
	EmailPort = int32(465)

	// EmailAddress 邮箱地址
	EmailAddress = "546247083@qq.com"

	// EmailName 邮箱名称
	EmailName = "546247083@qq.com"

	// EmailPass 邮箱密码
	EmailPass = "fhdwnwhjcieobdja"

	// UploadPath 上传路径
	UploadPath = "./upload/"

	// GoServerAddress GoServerAddress的地址
	GoServerAddress = "http://107.151.172.51:8883/"

	// FileServerAddress FileServerAddress的地址
	FileServerAddress = "http://107.151.172.51:8882/"

	// ChatServerAddress ChatServerAddress的地址
	ChatServerAddress = "http://107.151.172.51:8884/"

	// 读取的配置
	xmlConfig *XmlConfig

	// 错误
	err error
)

func init() {
	// 读取配置文件
	xmlConfig = NewXmlConfig()
	err = xmlConfig.LoadFromFile("config.xml")
	checkError(err, true)

	// 读取服务器配置端口
	RPCListenAddress, err = xmlConfig.String("root/BaseConfig/RPCListenAddress", "")
	checkError(err, false)

	WebListenAddress, err = xmlConfig.String("root/BaseConfig/WebListenAddress", "")
	checkError(err, false)

	// 读取数据库配置
	DBConnection, err = xmlConfig.String("root/DBConnection/WebServer", "")
	checkError(err, false)

	// 读取日志配置
	IsDebug, err = xmlConfig.Bool("root/LogConfig/IsDebug", "")
	checkError(err, false)

	LogPath, err = xmlConfig.String("root/LogConfig/LogPath", "")
	checkError(err, false)

	LogInfoFlag, err = xmlConfig.Bool("root/LogConfig/LogInfoFlag", "")
	checkError(err, false)

	LogDebugFlag, err = xmlConfig.Bool("root/LogConfig/LogDebugFlag", "")
	checkError(err, false)

	LogWarnFlag, err = xmlConfig.Bool("root/LogConfig/LogWarnFlag", "")
	checkError(err, false)

	LogErrorFlag, err = xmlConfig.Bool("root/LogConfig/LogErrorFlag", "")
	checkError(err, false)

	LogFatalFlag, err = xmlConfig.Bool("root/LogConfig/LogFatalFlag", "")
	checkError(err, false)

	// 读取网站配置
	WebMainPath, err = xmlConfig.String("root/WebConfig/WebMainPath", "")
	checkError(err, false)

	IndexPage, err = xmlConfig.String("root/WebConfig/IndexPage", "")
	checkError(err, false)

	Error404Page, err = xmlConfig.String("root/WebConfig/Error404Page", "")
	checkError(err, false)

	Error500Page, err = xmlConfig.String("root/WebConfig/Error500Page", "")
	checkError(err, false)

	// 读取服务器的网站配置
	WebSiteReferer, err = xmlConfig.String("root/WebSiteConfig/WebSiteReferer", "")
	checkError(err, false)

	PwdExpiredTime, err = xmlConfig.Int("root/WebSiteConfig/PwdExpiredTime", "")
	checkError(err, false)

	EmailHost, err = xmlConfig.String("root/WebSiteConfig/EmailHost", "")
	checkError(err, false)

	EmailPort, err = xmlConfig.Int32("root/WebSiteConfig/EmailPort", "")
	checkError(err, false)

	EmailAddress, err = xmlConfig.String("root/WebSiteConfig/EmailAddress", "")
	checkError(err, false)

	EmailName, err = xmlConfig.String("root/WebSiteConfig/EmailName", "")
	checkError(err, false)

	EmailPass, err = xmlConfig.String("root/WebSiteConfig/EmailPass", "")
	checkError(err, false)

	//读取服务器配置
	FileServerAddress, err = xmlConfig.String("root/ServerConfig/FileServerAddress", "")
	checkError(err, false)

	GoServerAddress, err = xmlConfig.String("root/ServerConfig/GoServerAddress", "")
	checkError(err, false)

	ChatServerAddress, err = xmlConfig.String("root/ServerConfig/ChatServerAddress", "")
	checkError(err, false)

	//读取文件配置
	UploadPath, err = xmlConfig.String("root/FileConfig/UploadPath", "")
	checkError(err, false)
}

// checkError 抛出错误
// 是否抛出错误
func checkError(err error, ifThrowError bool) {
	if err != nil && ifThrowError {
		panic(err)
	}
}
