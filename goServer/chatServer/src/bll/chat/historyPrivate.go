package chat

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/chatServer/src/dal"
	"xq.goproject.com/goServer/chatServer/src/model"
)

var (
	//私聊聊天记录
	historyPrivateMap = make(map[string][]*model.HistoryPrivate)
)

func init() {
	initTool.RegisterInitFunc(initHistoryPrivateData, initTool.I_NeedInit)
}

// 初始化数据
func initHistoryPrivateData() error {
	historyPrivateList, err := dal.HistoryPrivateDALObj.GetAllList()
	if err != nil {
		return err
	}

	//初始化数据
	for _, item := range historyPrivateList {
		if historyPrivateMap[item.SysUserID] == nil {
			historyPrivateMap[item.SysUserID] = make([]*model.HistoryPrivate, 0, 32)
		}

		historyPrivateMap[item.SysUserID] = append(historyPrivateMap[item.SysUserID], item)
	}

	return nil
}

// GetHistoryPrivateList 获取用户的私聊聊天记录
func GetHistoryPrivateList(userID string) []*model.HistoryPrivate {
	result := make([]*model.HistoryPrivate, 0, 32)
	resultTemp, exist := historyPrivateMap[userID]
	if exist {
		result = resultTemp
	}

	return result
}
