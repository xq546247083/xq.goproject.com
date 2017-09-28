package initDataTool

import (
	"fmt"

	"xq.goproject.com/commonTools/logTool"
)

var (
	// 初始化数据Func字典
	initDataFuncMap = make(map[InitDataType][]func())

	// 初始化数据类型
	initDataTypeList = [4]InitDataType{
		Config,
		DynamicConfig,
		Global,
		NeedInit,
	}
)

// RegisterFunc 注册初始化数据的方法
func RegisterFunc(initDataType InitDataType, registerFunc func()) {
	if _, exists := initDataFuncMap[initDataType]; !exists {
		initDataFuncMap[initDataType] = make([]func(), 0, 32)
	}

	initDataFuncMap[initDataType] = append(initDataFuncMap[initDataType], registerFunc)
	logTool.Log(logTool.Info, fmt.Sprintf("注册类型:%s，当前共有%d个注册", initDataType.ToString(), len(initDataFuncMap[initDataType])))
}
