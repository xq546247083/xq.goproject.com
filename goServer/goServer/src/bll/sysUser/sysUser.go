package sysUser

import (
	"time"

	"xq.goproject.com/commonTools/EncrpytTool"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
	"xq.goproject.com/goServer/goServerModel/src/consts"
)

var (
	//用户数据缓存
	sysUserMap = make(map[string]*model.SysUser)

	//用户发送邮箱缓存
	//key:邮箱
	sysUserEmailMap = make(map[string]*model.SysUserEmail)
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

//UpdatePwdExpiredTime 更新过期时间
func UpdatePwdExpiredTime(userNameOrEmail string) {
	sysUser := GetItemByUserNameOrEmail(userNameOrEmail)
	if sysUser != nil {
		//修改密码过期时间
		duration := time.Duration(int(time.Hour) * configTool.PwdExpiredTime)
		sysUser.PwdExpiredTime = sysUser.PwdExpiredTime.Add(duration)

		//保存数据
		dal.SysUserDALObj.SaveInfo(sysUser, nil)
	}
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

// 组装数据返回
func assembleToClient(sysUser *model.SysUser) map[string]interface{} {
	clientInfo := make(map[string]interface{})

	clientInfo[consts.UserName] = sysUser.UserName
	clientInfo[consts.FullName] = sysUser.FullName
	clientInfo[consts.Sex] = sysUser.Sex
	clientInfo[consts.Phone] = sysUser.Phone
	clientInfo[consts.Email] = sysUser.Email
	clientInfo[consts.LastLoginTime] = sysUser.LastLoginTime
	clientInfo[consts.LastLoginIP] = sysUser.LastLoginIP
	clientInfo[consts.LoginCount] = sysUser.LoginCount
	clientInfo[consts.Status] = sysUser.Status
	clientInfo[consts.CreateTime] = sysUser.CreateTime
	clientInfo[consts.PwdExpiredTime] = sysUser.PwdExpiredTime.UnixNano() / 1e6

	return clientInfo
}

// 组装数据返回
func assembleToClientAllUser() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})
	for _, sysUser := range sysUserMap {
		clientInfo := make(map[string]interface{})

		clientInfo[consts.UserID] = sysUser.UserID
		clientInfo[consts.UserName] = sysUser.UserName
		clientInfo[consts.FullName] = sysUser.FullName
		clientInfo[consts.Sex] = sysUser.Sex
		clientInfo[consts.Phone] = sysUser.Phone
		clientInfo[consts.Email] = sysUser.Email
		clientInfo[consts.LastLoginTime] = sysUser.LastLoginTime
		clientInfo[consts.LastLoginIP] = sysUser.LastLoginIP
		clientInfo[consts.LoginCount] = sysUser.LoginCount
		clientInfo[consts.Status] = sysUser.Status
		clientInfo[consts.CreateTime] = sysUser.CreateTime
		clientInfo[consts.PwdExpiredTime] = sysUser.PwdExpiredTime

		result[sysUser.UserName] = clientInfo
	}

	return result
}
