package initTool

import (
	"fmt"

	"xq.goproject.com/commonTools/logTool"
)

var (
	// 初始化数据Func字典
	initDataFuncMap = make(map[InitDataType][]func() error)

	// 初始化数据类型
	initDataTypeList = [4]InitDataType{
		I_Config,
		I_DynamicConfig,
		I_Global,
		I_NeedInit,
	}
)

// RegisterInitFunc 注册初始化数据的方法
func RegisterInitFunc(registerFunc func() error, initDataType InitDataType) {
	if _, exists := initDataFuncMap[initDataType]; !exists {
		initDataFuncMap[initDataType] = make([]func() error, 0, 32)
	}

	initDataFuncMap[initDataType] = append(initDataFuncMap[initDataType], registerFunc)
	logTool.Log(logTool.Info, fmt.Sprintf("初始化数据方法中，注册类型:%s，当前共有%d个注册", initDataType.ToString(), len(initDataFuncMap[initDataType])))
}
