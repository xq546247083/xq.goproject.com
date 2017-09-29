package initTool

import (
	"fmt"

	"xq.goproject.com/commonTools/fmtTool"
	"xq.goproject.com/commonTools/logTool"
)

//checkData 检测数据
func checkData() {
	//错误消息
	errMsg := make([]string, 0, 32)

	//循环字典，调用检测数据的方法
	for _, checkDataType := range checkDataTypeList {
		if items, exists := checkDataFuncMap[checkDataType]; exists {
			logTool.Log(logTool.Info, fmt.Sprintf("检测【%s】数据开始", checkDataType.ToString()))

			//构造错误数据
			errMsg = errMsg[:0]
			errMsg = append(errMsg, fmt.Sprintf("检测【%s】数据出错,err:", checkDataType.ToString()))

			//循环检测数据
			for _, itemFunc := range items {
				errs := itemFunc()
				for _, errItem := range errs {
					errMsg = append(errMsg, errItem.Error())
				}
			}

			if len(errMsg) >= 2 {
				fmtTool.Println(errMsg...)
				logTool.Log(logTool.Error, errMsg...)
				panic(fmt.Errorf("检测【%s】数据失败！", checkDataType.ToString()))
			} else {
				logTool.Log(logTool.Info, fmt.Sprintf("检测【%s】数据结束", checkDataType.ToString()))
			}
		}
	}
}
