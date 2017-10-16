package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// historyPrivateDAL 私聊信息记录dal
type historyPrivateDAL struct{}

var (
	// HistoryPrivateDALObj dal数据对象
	HistoryPrivateDALObj = new(historyPrivateDAL)

	// DALName 连接对象名
	historyPrivateDALName = "HistoryPrivateDALObj"
)

// GetAllList 获取数据
func (thisObj *historyPrivateDAL) GetAllList() (historyPrivateList []*model.HistoryPrivate, err error) {
	if err = DB.Find(&historyPrivateList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", historyPrivateDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *historyPrivateDAL) SaveInfo(historyPrivate *model.HistoryPrivate, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(historyPrivate).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", historyPrivateDALName))
		return
	}

	return nil
}
