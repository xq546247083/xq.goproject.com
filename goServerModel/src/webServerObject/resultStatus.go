package webServerObject

//ResultStatus 服务端请求，中心服务响应结果的状态对象，成功是0，非成功以负数来表示
type ResultStatus int

//ToString  返回响应状态枚举值对应的描述信息字符串
func (rs ResultStatus) ToString() string {
	return status[rs*-1]
}

// 定义所有的响应结果的状态枚举值，此种实现方式是GO语言的标准方式
const (
	// 成功
	Success ResultStatus = -1 * iota

	// 数据错误
	DataError

	// API数据错误
	APIDataError

	// 客户端数据错误
	ClientDataError

	// 命令类型未定义
	CommandTypeNotDefined

	// 签名错误
	SignError

	// 尚未登陆
	NoLogin

	// 不在公会中
	NotInUnion

	// 未找到目标
	NotFoundTarget

	// 不能给自己发消息
	CantSendMessageToSelf

	// 玩家不存在
	PlayerNotExist

	// 玩家被封号
	PlayerIsForbidden

	// 玩家被禁言
	PlayerIsInSilent

	// 只支持POST
	OnlySupportPOST

	// API未定义
	APINotDefined

	// 在另一台设备上登录
	LoginOnAnotherDevice

	// 名称错误
	NameError

	// 公会Id错误
	UnionIDError

	// 含有屏蔽词语
	ContainForbiddenWord

	//参数为空
	ParamIsEmpty

	//参数不匹配
	ParamNotMatch

	// 服务器组不存在
	ServerGroupNotExist

	// 不能发送跨服消息
	CantSendCrossServerMessage

	// 目标玩家不在线
	TargetPlayerOffline

	// 等级错误
	LvError

	// Vip等级错误
	VipError

	// 发送消息太快
	SendMessageTooFast

	// 等级不足，系统未开放
	LvIsNotEnough

	// 重复次数太多
	RepeatTooMuch

	//参数无效
	ParamInValid

	//没有目标方法
	NoTargetMethod

	//方法返回值无效
	ReturnValueIsValid

	//保存文件失败
	SaveFileFail
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var status = []string{
	"Success",
	"DataError",
	"APIDataError",
	"ClientDataError",
	"CommandTypeNotDefined",
	"SignError",
	"NoLogin",
	"NotInUnion",
	"NotFoundTarget",
	"CantSendMessageToSelf",
	"PlayerNotExist",
	"PlayerIsForbidden",
	"PlayerIsInSilent",
	"OnlySupportPOST",
	"APINotDefined",
	"LoginOnAnotherDevice",
	"NameError",
	"UnionIdError",
	"ContainForbiddenWord",
	"ParamIsEmpty",
	"ParamNotMatch",
	"ServerGroupNotExist",
	"CantSendCrossServerMessage",
	"TargetPlayerOffline",
	"LvError",
	"VipError",
	"SendMessageTooFast",
	"LvIsNotEnough",
	"RepeatTooMuch",
	"ParamInValid",
	"NoTargetMethod",
	"ReturnValueIsValid",
}
