package initTool

import (
	"fmt"

	"xq.goproject.com/commonTools/logTool"
)

//initData 初始化数据
func initData() {
	//循环字典，调用初始化数据的方法
	for _, initDataType := range initDataTypeList {
		if items, exists := initDataFuncMap[initDataType]; exists {
			logTool.Log(logTool.Info, fmt.Sprintf("初始化【%s】数据开始", initDataType.ToString()))
			for _, itemFunc := range items {
				//数据初始化失败，抛出错误
				if err := itemFunc(); err != nil {
					logTool.Log(logTool.Error, fmt.Sprintf("初始化【%s】数据出错,err:%s", initDataType.ToString(), err))
					panic(err)
				}
			}

			logTool.Log(logTool.Info, fmt.Sprintf("初始化【%s】数据结束", initDataType.ToString()))
		}
	}
}
