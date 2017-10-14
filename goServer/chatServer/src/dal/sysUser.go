package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// sysUserDAL 用户dal
type sysUserDAL struct{}

var (
	// SysUserDALObj dal数据对象
	SysUserDALObj = new(sysUserDAL)

	// sysUserDALName 连接对象名
	sysUserDALName = "SysUserDALObj"
)

// GetAllList 获取数据
func (thisObj *sysUserDAL) GetAllList() (sysUserList []*model.SysUser, err error) {
	if err = DB.Find(&sysUserList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", sysUserDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *sysUserDAL) SaveInfo(sysUser *model.SysUser, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(sysUser).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", sysUserDALName))
		return
	}

	return nil
}
