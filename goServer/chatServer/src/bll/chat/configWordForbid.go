package chat

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/chatServer/src/dal"
	"xq.goproject.com/goServer/chatServer/src/model"
)

var (
	// configWordForbidList 禁止单词列表
	configWordForbidList = make([]*model.ConfigWordForbid, 0, 32)
)

func init() {
	initTool.RegisterInitFunc(initConfigWordForbidData, initTool.I_NeedInit)
}

// 初始化数据
func initConfigWordForbidData() error {
	var err error
	if configWordForbidList, err = dal.ConfigWordForbidDALObj.GetAllList(); err != nil {
		return err
	}

	return nil
}
