package initTool

import (
	"fmt"

	"xq.goproject.com/commonTools/logTool"
)

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
	logTool.Log(logTool.Info, fmt.Sprintf("检测数据方法中，注册类型:%s，当前共有%d个注册", initDataType.ToString(), len(checkDataFuncMap[initDataType])))
}
