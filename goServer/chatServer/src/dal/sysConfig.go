package dal

import (
	"fmt"

	"xq.goproject.com/Vendor/github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// sysConfigDAL 系统配置dal
type sysConfigDAL struct{}

var (
	// SysConfigDALObj dal数据对象
	SysConfigDALObj = new(sysConfigDAL)

	// DALName 连接对象名
	sysConfigDALName = "SysConfigDALObj"
)

// GetAllList 获取数据
func (thisObj *sysConfigDAL) GetAllList() (sysConfigList []*model.SysConfig, err error) {
	if err = DB.Find(&sysConfigList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", sysConfigDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *sysConfigDAL) SaveInfo(sysConfig *model.SysConfig, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(sysConfig).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", sysConfigDALName))
		return
	}

	return nil
}
