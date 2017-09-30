package webServerObject

//ResultStatus 服务端请求，中心服务响应结果的状态对象，成功是0，非成功以负数来表示
type ResultStatus int

//ToString  返回响应状态枚举值对应的描述信息字符串
func (rs ResultStatus) ToString() string {
	return status[rs*-1]
}

//ToDescription  返回描述信息
func (rs ResultStatus) ToDescription() string {
	return description[rs*-1]
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

	// 密码已过期
	LoginIsOver

	// 密码错误
	PwdError

	// 用户不存在
	UserIsNotExist

	// 用户名已被注册
	UserNameIsExist

	// 用户名必须以字母开头
	UserNameMustBeginWithLetter

	// 用户名只能由字母和数字构成
	UserNameMustBeLetterOrNum

	// 密码不能为空
	UserPasswordCanBeNotEmpty

	// 邮箱不能为空
	EmailCanBeNotEmpty

	// 邮箱格式错误
	EmailStyleIsError

	// 邮箱已被注册
	EmailAlreadyExist

	// 用户名不能为空
	UserNameCantBeEmpty

	// 电话号码格式错误
	PhoneStyleIsError

	// 请输入密码
	PlsEnterPassword

	// 请输入验证码
	PlsEnterIdentifyCode

	// 该邮箱还未发送验证码
	IdentifyCodeNoThisEmail

	// 验证码错误
	IdentifyCodeIsError

	// 发送邮件失败，请检查邮箱
	SendEmailFail

	// 发送邮件过快，请稍后重试
	SendEmailIsFast

	// 该邮箱未注册
	EmailIsNotRegister
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
	"UnionIDError",
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
	"SaveFileFail",
	"LoginIsOver",
	"PwdError",
	"UserIsNotExist",
	"UserNameIsExist",
	"UserNameMustBeginWithLetter",
	"UserNameMustBeLetterOrNum",
	"UserPasswordCanBeNotEmpty",
	"EmailCanBeNotEmpty",
	"EmailStyleIsError",
	"EmailAlreadyExist",
	"UserNameCantBeEmpty",
	"PhoneStyleIsError",
	"PlsEnterPassword",
	"PlsEnterIdentifyCode",
	"IdentifyCodeNoThisEmail",
	"IdentifyCodeIsError",
	"SendEmailFail",
	"SendEmailIsFast",
	"EmailIsNotRegister",
}

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var description = []string{
	"成功",
	"数据错误",
	"API数据错误",
	"客户端数据错误",
	"命令类型未定义",
	"签名错误",
	"尚未登陆",
	"不在公会中",
	"未找到目标",
	"不能给自己发消息",
	"玩家不存在",
	"玩家被封号",
	"玩家被禁言",
	"只支持POST",
	"API未定义",
	"在另一台设备上登录",
	"名称错误",
	"公会Id错误",
	"含有屏蔽词语",
	"数为空",
	"数不匹配",
	"服务器组不存在",
	"不能发送跨服消息",
	"目标玩家不在线",
	"等级错误",
	"Vip等级错误",
	"发送消息太快",
	"等级不足，系统未开放",
	"重复次数太多",
	"数无效",
	"有目标方法",
	"法返回值无效",
	"存文件失败",
	"密码已过期",
	"密码错误",
	"用户不存在",
	"用户名已被注册",
	"用户名必须以字母开头",
	"用户名只能由字母和数字构成",
	"密码不能为空",
	"邮箱不能为空",
	"邮箱格式错误",
	"邮箱已被注册",
	"用户名不能为空",
	"电话号码格式错误",
	"请输入密码",
	"请输入验证码",
	"该邮箱还未发送验证码",
	"验证码错误",
	"发送邮件失败，请检查邮箱",
	"发送邮件过快，请稍后重试",
	"该邮箱未注册",
}
