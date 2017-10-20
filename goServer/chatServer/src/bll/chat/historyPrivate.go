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
		if historyPrivateMap[item.SysUserName] == nil {
			historyPrivateMap[item.SysUserName] = make([]*model.HistoryPrivate, 0, 32)
		}

		historyPrivateMap[item.SysUserName] = append(historyPrivateMap[item.SysUserName], item)
	}

	return nil
}

// getUnSendHistoryPrivateList 获取用户的离线私聊聊天记录
func getUnSendHistoryPrivateList(userName string) []*model.HistoryPrivate {
	result := make([]*model.HistoryPrivate, 0, 32)

	//获取用户聊天记录
	resultTemp, exist := historyPrivateMap[userName]
	if exist {
		for _, historyPrivate := range resultTemp {
			if !historyPrivate.IsSend {
				result = append(result, historyPrivate)
			}
		}
	}

	return result
}

// savetHistoryPrivate 保存私聊聊天记录
func savetHistoryPrivate(histtoryPrivate *model.HistoryPrivate) error {
	return dal.HistoryPrivateDALObj.SaveInfo(histtoryPrivate, nil)
}
