package initTool

//InitDataType 初始化数据的类型
type InitDataType int32

//ToString 返回字符串
func (initDataType InitDataType) ToString() string {
	return initDataTypeString[initDataType-1]
}

const (
	//I_Config 基础配置
	I_Config InitDataType = 1 + iota

	//I_DynamicConfig 动态配置
	I_DynamicConfig

	//I_Global 全局数据
	I_Global

	//I_NeedInit 需要初始化的数据
	I_NeedInit
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var initDataTypeString = []string{
	"Config",
	"DynamicConfig",
	"Global",
	"NeedInit",
}
