package sysUser

import (
	"encoding/json"
	"sync"
	"time"

	"xq.goproject.com/commonTools/EncrpytTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/chatServer/src/model"
	"xq.goproject.com/goServer/chatServer/src/webClient"
)

var (
	//用户数据缓存
	sysUserMap = make(map[string]*model.SysUser)

	//锁
	mutex sync.RWMutex
)

func init() {
	initTool.RegisterInitFunc(initSysUserData, initTool.I_NeedInit)
}

// initSysUserData 初始化用户数据
func initSysUserData() error {
	// 获取用户数据
	responseObj, err := webClient.PostDataToGoServer(webClient.GetAllUser, []interface{}{}, false)
	if err != nil {
		return err
	}

	//反序列化字典为byte
	dataByte, err2 := json.Marshal(responseObj.Data)
	if err2 != nil {
		return err2
	}

	//再序列化为对象
	err3 := json.Unmarshal(dataByte, &sysUserMap)
	if err3 != nil {
		return err3
	}

	if sysUserMap == nil {
		sysUserMap = make(map[string]*model.SysUser)
	}

	return nil
}

// GetAllSysUser 获取所有用户
func GetAllSysUser() map[string]*model.SysUser {
	return sysUserMap
}

// 添加或者更新用户用户
func updateSysUser(sysUser *model.SysUser) {
	mutex.Lock()
	defer mutex.Unlock()

	sysUserMap[sysUser.UserName] = sysUser
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
