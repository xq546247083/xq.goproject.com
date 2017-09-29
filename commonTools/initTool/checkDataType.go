package initTool

//CheckDataType 初始化数据的类型
type CheckDataType int32

//ToString 返回字符串
func (checkDataType CheckDataType) ToString() string {
	return checkDataTypeString[checkDataType-1]
}

const (
	//C_Config 基础配置
	C_Config CheckDataType = 1 + iota

	//C_DynamicConfig 动态配置
	C_DynamicConfig
)

// 定义所有的响应结果的状态值所对应的字符串描述信息，如果要增加状态枚举，则此处也要相应地增加
var checkDataTypeString = []string{
	"Config",
	"DynamicConfig",
}
