package initDataTool

import (
	"fmt"

	"xq.goproject.com/commonTools/logTool"
)

//InitData 初始化数据
func InitData() {
	//循环字典，调用初始化数据的方法
	for _, initDataType := range initDataTypeList {
		if items, exists := initDataFuncMap[initDataType]; exists {
			logTool.Log(logTool.Info, fmt.Sprintf("初始化【%s】数据开始", initDataType.ToString()))
			for _, item := range items {
				item()
			}

			logTool.Log(logTool.Info, fmt.Sprintf("初始化【%s】数据结束", initDataType.ToString()))
		}
	}
}
