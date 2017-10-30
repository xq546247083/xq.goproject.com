package dal

import (
	"fmt"

	"xq.goproject.com/Vendor/github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// configWordSensitiveDAL 敏感单词dal
type configWordSensitiveDAL struct{}

var (
	// ConfigWordSensitiveDALObj dal数据对象
	ConfigWordSensitiveDALObj = new(configWordSensitiveDAL)

	// DALName 连接对象名
	configWordSensitiveDALName = "ConfigWordSensitiveDALObj"
)

// GetAllList 获取数据
func (thisObj *configWordSensitiveDAL) GetAllList() (configWordSensitiveList []*model.ConfigWordSensitive, err error) {
	if err = DB.Find(&configWordSensitiveList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", configWordSensitiveDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *configWordSensitiveDAL) SaveInfo(configWordSensitive *model.ConfigWordSensitive, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(configWordSensitive).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", configWordSensitiveDALName))
		return
	}

	return nil
}
