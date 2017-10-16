package chat

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/chatServer/src/dal"
	"xq.goproject.com/goServer/chatServer/src/model"
)

var (
	// configWordSensitiveList 敏感单词列表
	configWordSensitiveList = make([]*model.ConfigWordSensitive, 0, 32)
)

func init() {
	initTool.RegisterInitFunc(initConfigWordSensitiveData, initTool.I_NeedInit)
}

// 初始化数据
func initConfigWordSensitiveData() error {
	var err error
	if configWordSensitiveList, err = dal.ConfigWordSensitiveDALObj.GetAllList(); err != nil {
		return err
	}

	return nil
}
