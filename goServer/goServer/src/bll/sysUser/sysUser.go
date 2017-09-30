package sysUser

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

var (
	sysUserMap = make(map[string]*model.SysUser)
)

func init() {
	initTool.RegisterInitFunc(initData, initTool.I_NeedInit)
}

// 初始化数据
func initData() error {
	sysUserList, err := dal.SysUserDALObj.GetAllList()
	if err != nil {
		return err
	}

	for _, item := range sysUserList {
		sysUserMap[item.UserID] = item
	}

	return nil
}
