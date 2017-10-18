package chat

import (
	"xq.goproject.com/goServer/chatServer/src/dal"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// GetHistoryWorldList 获取世界聊天记录
func GetHistoryWorldList(num int32) []*model.HistoryWorld {
	historyWorldList, err := dal.HistoryWorldDALObj.GetAllList()
	if err == nil {
		return make([]*model.HistoryWorld, 0, 32)
	}

	length := int32(len(historyWorldList))

	// InsertHistoryWorld获取截取数据的位置
	front := length - num
	if length < num {
		front = 0
	}

	return historyWorldList[front:length]
}

// savetHistoryWorld 保存世界聊天记录
func savetHistoryWorld(histtoryWorld *model.HistoryWorld) error {
	return dal.HistoryWorldDALObj.SaveInfo(histtoryWorld, nil)
}
