package userBLL

import (
	"xq.goproject.com/commonTools/initTool"
)

func init() {
	initTool.RegisterInitFunc(initData, initTool.I_Config)
	initTool.RegisterCheckFunc(checkData, initTool.C_Config)
}

func initData() error {
	return nil
}

func checkData() []error {
	return nil
}
