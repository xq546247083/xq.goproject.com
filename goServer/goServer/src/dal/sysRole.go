package dal

import (
	"fmt"

	"xq.goproject.com/goServer/goServer/src/model"
)

// sysRoleDAL 角色dal
type sysRoleDAL struct{}

var (
	// SysRoleDALObj dal数据对象
	SysRoleDALObj = new(sysRoleDAL)

	// DALName 连接对象名
	sysRoleDALName = "SysRoleDALObj"
)

// GetAllList 获取数据
func (thisObj *sysRoleDAL) GetAllList() (sysRoleList []*model.SysRole, err error) {
	if err = DB.Find(&sysRoleList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", sysRoleDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *sysRoleDAL) SaveInfo(sysRole *model.SysRole) (err error) {
	if err = DB.Save(sysRole).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", sysRoleDALName))
		return
	}

	return nil
}
