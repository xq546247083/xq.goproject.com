package dal

import (
	"fmt"

	"xq.goproject.com/Vendor/github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// historyWorldDAL 世界聊天信息记录dal
type historyWorldDAL struct{}

var (
	// HistoryWorldDALObj dal数据对象
	HistoryWorldDALObj = new(historyWorldDAL)

	// DALName 连接对象名
	historyWorldDALName = "HistoryWorldDALObj"
)

// GetAllList 获取数据
func (thisObj *historyWorldDAL) GetAllList() (historyWorldList []*model.HistoryWorld, err error) {
	if err = DB.Find(&historyWorldList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", historyWorldDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *historyWorldDAL) SaveInfo(historyWorld *model.HistoryWorld, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(historyWorld).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", historyWorldDALName))
		return
	}

	return nil
}
