package sysUser

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/stringTool"
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

// GetItemByUserNameOrEmail 获取用户通过用户名或者邮箱
func GetItemByUserNameOrEmail(userNameOrEmail string) (sysUser *model.SysUser) {
	//通过用户名或者邮箱获取用户
	if stringTool.IsEmail(userNameOrEmail) {
		for _, sysUser := range sysUserMap {
			if sysUser.Email == userNameOrEmail {
				return sysUser
			}
		}
	} else {
		for _, sysUser := range sysUserMap {
			if sysUser.UserName == userNameOrEmail {
				return sysUser
			}
		}
	}

	return nil
}
