package dal

import (
	"fmt"

	"xq.goproject.com/goServer/goServer/src/model"
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
func (thisObj *sysUserDAL) SaveInfo(sysUser *model.SysUser) (err error) {
	if err = DB.Save(sysUser).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", sysUserDALName))
		return
	}

	return nil
}
