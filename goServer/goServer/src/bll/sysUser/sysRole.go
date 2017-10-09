package sysUser

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

var (
	//角色缓存数据
	sysRoleMap = make(map[int32]*model.SysRole)
)

func init() {
	initTool.RegisterInitFunc(initSysRoleData, initTool.I_NeedInit)
}

// 初始化数据
func initSysRoleData() error {
	sysRoleList, err := dal.SysRoleDALObj.GetAllList()
	if err != nil {
		return err
	}

	for _, item := range sysRoleList {
		sysRoleMap[item.RoleID] = item
	}

	return nil
}
