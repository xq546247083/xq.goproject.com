package initTool

var (

	// 检测数据Func字典
	checkDataFuncMap = make(map[CheckDataType][]func() []error)

	// 检测数据类型
	checkDataTypeList = [2]CheckDataType{
		C_Config,
		C_DynamicConfig,
	}
)

// RegisterCheckFunc 注册检测数据的方法
func RegisterCheckFunc(registerFunc func() []error, initDataType CheckDataType) {
	if _, exists := checkDataFuncMap[initDataType]; !exists {
		checkDataFuncMap[initDataType] = make([]func() []error, 0, 32)
	}

	checkDataFuncMap[initDataType] = append(checkDataFuncMap[initDataType], registerFunc)
}
