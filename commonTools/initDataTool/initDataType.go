package initDataTool

//InitDataType 初始化数据的类型
type InitDataType int32

//ToString 返回字符串
func (initDataType InitDataType) ToString() string {
	return initDataTypeString[initDataType-1]
}

const (
	//Config 基础配置
	Config InitDataType = 1 + iota

	//DynamicConfig 动态配置
	DynamicConfig

	//Global 全局数据
	Global

	//NeedInit 需要初始化的数据
	NeedInit
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var initDataTypeString = []string{
	"Config",
	"DynamicConfig",
	"Global",
	"NeedInit",
}
