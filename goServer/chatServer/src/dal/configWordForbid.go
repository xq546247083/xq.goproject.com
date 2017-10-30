package dal

import (
	"fmt"

	"xq.goproject.com/Vendor/github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// configWordForbidDAL 禁止单词dal
type configWordForbidDAL struct{}

var (
	// ConfigWordForbidDALObj dal数据对象
	ConfigWordForbidDALObj = new(configWordForbidDAL)

	// DALName 连接对象名
	configWordForbidDALName = "ConfigWordForbidDALObj"
)

// GetAllList 获取数据
func (thisObj *configWordForbidDAL) GetAllList() (configWordForbidList []*model.ConfigWordForbid, err error) {
	if err = DB.Find(&configWordForbidList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", configWordForbidDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *configWordForbidDAL) SaveInfo(configWordForbid *model.ConfigWordForbid, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(configWordForbid).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", configWordForbidDALName))
		return
	}

	return nil
}
