package sysUser

import (
	"sync"
	"time"

	"xq.goproject.com/commonTools/EncrpytTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/chatServer/src/model"
)

var (
	//用户数据缓存
	sysUserMap = make(map[string]*model.SysUser)

	//锁
	mutex sync.RWMutex
)

// 添加或者更新用户用户
func updateSysUser(sysUser *model.SysUser) {
	mutex.Lock()
	defer mutex.Unlock()

	sysUserMap[sysUser.UserID] = sysUser
}

// GetItemByUserNameOrEmail 获取用户通过用户名或者邮箱
func GetItemByUserNameOrEmail(userNameOrEmail string) (sysUser *model.SysUser) {
	mutex.RLock()
	defer mutex.RUnlock()

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

//CheckPwdExpiredTime 检测密码过期时间
func CheckPwdExpiredTime(userNameOrEmail string) bool {
	sysUser := GetItemByUserNameOrEmail(userNameOrEmail)
	if sysUser != nil {
		return sysUser.PwdExpiredTime.Unix() < time.Now().Unix()
	}

	return false
}

//GetUserToken 获取用户token
func GetUserToken(userNameOrEmail string) string {
	sysUser := GetItemByUserNameOrEmail(userNameOrEmail)
	if sysUser != nil {
		timeStamp := sysUser.LastLoginTime.Format("2000-01-04 01:01:01")
		result := EncrpytTool.Encrypt(sysUser.UserName + "!A%HS*I^" + timeStamp)
		return result
	}

	return ""
}
