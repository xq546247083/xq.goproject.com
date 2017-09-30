package dal

import (
	"fmt"

	"xq.goproject.com/goServer/goServer/src/model"
)

// sysMenuDAL 菜单dal
type sysMenuDAL struct{}

var (
	// SysMenuDALObj dal数据对象
	SysMenuDALObj = new(sysMenuDAL)

	// sysMenuDALName 连接对象名
	sysMenuDALName = "SysMenuDALObj"
)

// GetAllList 获取数据
func (thisObj *sysMenuDAL) GetAllList() (sysMenuList []*model.SysMenu, err error) {
	if err = DB.Find(&sysMenuList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", sysMenuDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *sysMenuDAL) SaveInfo(sysMenu *model.SysMenu) (err error) {
	if err = DB.Save(sysMenu).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", sysMenuDALName))
		return
	}

	return nil
}
