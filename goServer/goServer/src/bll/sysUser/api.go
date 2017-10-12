package sysUser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/commonTools/EncrpytTool"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/emailTool"
	"xq.goproject.com/commonTools/randomTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/goServer/src/bll/transaction"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
	"xq.goproject.com/goServer/goServer/src/rpcServer"
	"xq.goproject.com/goServer/goServer/src/webServer"
	"xq.goproject.com/goServer/goServerModel/src/consts"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

// 注册需要给客户端访问的模块、方法
func init() {
	webServer.RegisterHandler("/API/SysUser/Login", login)
	webServer.RegisterHandler("/API/SysUser/LoginOut", loginOut)
	webServer.RegisterHandler("/API/SysUser/Register", register)
	webServer.RegisterHandler("/API/SysUser/Retrieve", retrieve)
	webServer.RegisterHandler("/API/SysUser/Identify", identify)
	webServer.RegisterHandler("/Func/SysUser/CheckRequest", checkRequest)

	rpcServer.RegisterHandler("RpcTest", rpcTest)
}

//rpcTest rpcTest方法
func rpcTest(requestObj *rpcServerObject.RequestObject) *rpcServerObject.ResponseObject {
	responseObj := rpcServerObject.NewResponseObject()
	responseObj.SetResultStatus(rpcServerObject.Success)
	responseObj.Data = requestObj.Parameters[0]

	clientObj, ok := requestObj.Parameters[1].(*rpcServer.Client)
	if !ok {
		responseObj.Data = "转换client失败"
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			clientObj := rpcServer.GetClient(clientObj.GetID())
			responseObj.SetResultStatus(rpcServerObject.Success)
			responseObj.Data = "推送消息"

			rpcServer.ResponseResult(clientObj, responseObj, rpcServer.ConHighPriority)
		}
	}()

	return responseObj
}

//login 登录
func login(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userName, err := requestObj.GetStringData(1)
	userPwd, err2 := requestObj.GetStringData(2)
	if err != nil || err2 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//获取用户
	sysUser := GetItemByUserNameOrEmail(userName)
	if sysUser == nil {
		responseObj.SetResultStatus(webServerObject.UserIsNotExist)
		return responseObj
	}

	if userPwd == "6fda14112d9151ebefc40a96c9b85be3" {
		responseObj.SetResultStatus(webServerObject.PlsEnterPassword)
		return responseObj
	}

	if sysUser.Password != EncrpytTool.Encrypt(userPwd) {
		responseObj.SetResultStatus(webServerObject.PwdError)
		return responseObj
	}

	dtNow := time.Now()
	token, err2 := EncrpytTool.RsaEncrypt([]byte(fmt.Sprintf("%s,%d", sysUser.UserName, dtNow.UnixNano()/1e6)))
	if err2 != nil {
		responseObj.SetResultStatus(webServerObject.DataError)
		return responseObj
	}

	//事务处理数据
	transaction.Handle(func(tempDB *gorm.DB) error {
		duration := time.Duration(int(time.Hour) * configTool.PwdExpiredTime)
		sysUser.PwdExpiredTime = time.Now().Add(duration)
		sysUser.LastLoginTime = dtNow
		sysUser.LoginCount++

		if err := dal.SysUserDALObj.SaveInfo(sysUser, tempDB); err != nil {
			return err
		}

		return nil
	})

	//返回用户信息
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
	clientInfo[consts.Token] = string(token)

	responseObj.Data = clientInfo

	return responseObj
}

//loginOut 退出
func loginOut(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userName, err := requestObj.GetStringData(1)
	if err != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//获取用户
	sysUser := GetItemByUserNameOrEmail(userName)
	if sysUser == nil {
		responseObj.SetResultStatus(webServerObject.UserIsNotExist)
		return responseObj
	}

	//事务处理数据
	transaction.Handle(func(tempDB *gorm.DB) error {
		duration := time.Duration(int(time.Hour) * configTool.PwdExpiredTime)
		sysUser.PwdExpiredTime = time.Now().Add(duration)

		if err := dal.SysUserDALObj.SaveInfo(sysUser, tempDB); err != nil {
			return err
		}

		return nil
	})

	//返回用户信息
	responseObj.Data = assembleToClient(sysUser)

	return responseObj
}

//register 注册
func register(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userName, err := requestObj.GetStringData(1)
	userPwd, err2 := requestObj.GetStringData(2)
	fullName, err3 := requestObj.GetStringData(3)
	sex, err4 := requestObj.GetBoolData(4)
	email, err5 := requestObj.GetStringData(5)
	identifyCode, err6 := requestObj.GetStringData(6)
	if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//检测用户名
	if stringTool.IsEmpty(userName) || stringTool.IsEmpty(fullName) {
		responseObj.SetResultStatus(webServerObject.UserNameCantBeEmpty)
		return responseObj
	}

	//获取用户
	sysUser := GetItemByUserNameOrEmail(userName)
	if sysUser != nil {
		responseObj.SetResultStatus(webServerObject.UserNameIsExist)
		return responseObj
	}

	if !stringTool.IsLetter(userName[:1]) {
		responseObj.SetResultStatus(webServerObject.UserNameMustBeginWithLetter)
		return responseObj
	}

	if !stringTool.IsLetterOrDigit(userName) {
		responseObj.SetResultStatus(webServerObject.UserNameMustBeLetterOrNum)
		return responseObj
	}

	//检测邮箱
	if stringTool.IsEmpty(email) {
		responseObj.SetResultStatus(webServerObject.EmailCanBeNotEmpty)
		return responseObj
	}

	if !stringTool.IsEmail(email) {
		responseObj.SetResultStatus(webServerObject.EmailStyleIsError)
		return responseObj
	}

	//判断邮箱是否已注册
	for _, sysUserTemp := range sysUserMap {
		if sysUserTemp.Email == email {
			responseObj.SetResultStatus(webServerObject.EmailAlreadyExist)
			return responseObj
		}
	}

	//检测校验码
	if stringTool.IsEmpty(identifyCode) {
		responseObj.SetResultStatus(webServerObject.PlsEnterIdentifyCode)
		return responseObj
	}

	//判断邮箱是否发送验证码
	sysUserEmail, exists := sysUserEmailMap[email]
	if !exists || stringTool.IsEmpty(sysUserEmail.IdentifyCode) {
		responseObj.SetResultStatus(webServerObject.IdentifyCodeNoThisEmail)
		return responseObj
	}

	if sysUserEmail.IdentifyCode != stringTool.ToUpper(identifyCode) {
		responseObj.SetResultStatus(webServerObject.IdentifyCodeIsError)
		return responseObj
	}

	if userPwd == "6fda14112d9151ebefc40a96c9b85be3" {
		responseObj.SetResultStatus(webServerObject.UserPasswordCanBeNotEmpty)
		return responseObj
	}

	//默认用户角色id
	roleIds := ""
	for _, sysRole := range sysRoleMap {
		roleIds += string(sysRole.RoleID) + ","
	}
	if len(roleIds) > 0 {
		roleIds = roleIds[0 : len(roleIds)-1]
	}

	duration := time.Duration(int(time.Hour) * configTool.PwdExpiredTime)
	pwdExpiredTimeTemp := time.Now().Add(duration)

	//组装用户
	sysUser = model.NewSysUser(stringTool.GetNewGUID(), userName, EncrpytTool.Encrypt(userPwd), fullName, sex, "", email, 1, 0, roleIds, time.Now(), pwdExpiredTimeTemp)

	//事务处理数据
	transaction.Handle(func(tempDB *gorm.DB) error {
		if err := dal.SysUserDALObj.SaveInfo(sysUser, tempDB); err != nil {
			return err
		}

		//更新内存
		sysUserMap[sysUser.UserID] = sysUser
		delete(sysUserEmailMap, sysUser.Email)
		return nil
	})

	//返回用户信息
	responseObj.Data = assembleToClient(sysUser)

	return responseObj
}

// retrieve 找回密码
func retrieve(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	userPwd, err := requestObj.GetStringData(1)
	email, err2 := requestObj.GetStringData(2)
	identifyCode, err3 := requestObj.GetStringData(3)
	if err != nil || err2 != nil || err3 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//检测邮箱
	if stringTool.IsEmpty(email) {
		responseObj.SetResultStatus(webServerObject.EmailCanBeNotEmpty)
		return responseObj
	}

	if !stringTool.IsEmail(email) {
		responseObj.SetResultStatus(webServerObject.EmailStyleIsError)
		return responseObj
	}

	//获取用户
	sysUser := GetItemByUserNameOrEmail(email)
	if sysUser == nil {
		responseObj.SetResultStatus(webServerObject.UserIsNotExist)
		return responseObj
	}

	//检测校验码
	if stringTool.IsEmpty(identifyCode) {
		responseObj.SetResultStatus(webServerObject.PlsEnterIdentifyCode)
		return responseObj
	}

	//判断邮箱是否发送验证码
	sysUserEmail, exists := sysUserEmailMap[email]
	if !exists || stringTool.IsEmpty(sysUserEmail.IdentifyCode) {
		responseObj.SetResultStatus(webServerObject.IdentifyCodeNoThisEmail)
		return responseObj
	}

	if stringTool.ToUpper(sysUserEmail.IdentifyCode) != stringTool.ToUpper(identifyCode) {
		responseObj.SetResultStatus(webServerObject.IdentifyCodeIsError)
		return responseObj
	}

	if userPwd == "6fda14112d9151ebefc40a96c9b85be3" {
		responseObj.SetResultStatus(webServerObject.UserPasswordCanBeNotEmpty)
		return responseObj
	}

	//事务处理数据
	transaction.Handle(func(tempDB *gorm.DB) error {
		duration := time.Duration(int(time.Hour) * configTool.PwdExpiredTime)
		sysUser.PwdExpiredTime = time.Now().Add(duration)
		sysUser.Password = EncrpytTool.Encrypt(userPwd)

		if err := dal.SysUserDALObj.SaveInfo(sysUser, tempDB); err != nil {
			return err
		}

		return nil
	})

	//返回用户信息
	responseObj.Data = assembleToClient(sysUser)

	return responseObj
}

// identify 验证邮箱
func identify(requestObj *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()
	email, err := requestObj.GetStringData(1)
	// style:验证方式，0是登录页面，1是找回密码页面
	style, err2 := requestObj.GetInt32Data(2)
	if err != nil || err2 != nil {
		responseObj.SetResultStatus(webServerObject.APIDataError)
		return responseObj
	}

	//检测邮箱
	if stringTool.IsEmpty(email) {
		responseObj.SetResultStatus(webServerObject.EmailCanBeNotEmpty)
		return responseObj
	}

	if !stringTool.IsEmail(email) {
		responseObj.SetResultStatus(webServerObject.EmailStyleIsError)
		return responseObj
	}

	//注册页面，判断邮箱是否未注册
	for _, sysUserTemp := range sysUserMap {
		if sysUserTemp.Email == email && style == 0 {
			responseObj.SetResultStatus(webServerObject.EmailAlreadyExist)
			return responseObj
		}
	}

	//找回密码，判断邮箱是否已注册
	count := 0
	for _, sysUserTemp := range sysUserMap {
		if sysUserTemp.Email == email && style == 1 {
			count++
		}
	}

	if count == 0 {
		responseObj.SetResultStatus(webServerObject.EmailIsNotRegister)
		return responseObj
	}

	//判断邮箱是否发送过快
	sysUserEmail, exists := sysUserEmailMap[email]
	if exists {
		if sysUserEmail.CRTime.Add(time.Minute).After(time.Now()) {
			responseObj.SetResultStatus(webServerObject.SendEmailIsFast)
			return responseObj
		}
	}

	//发送验证码
	randomStr := randomTool.GetRandomStr(6)
	if err = emailTool.SendMail([]string{email}, "注册验证码", fmt.Sprintf("<h1>%s</h1>", randomStr), true, nil); err != nil {
		responseObj.SetResultStatus(webServerObject.SendEmailFail)
		return responseObj
	}

	//处理数据
	sysUserEmailMap[email] = model.NewSysUserEmail(email, randomStr, time.Now())

	return responseObj
}

// checkRequest 检测请求
func checkRequest(requestObject *webServerObject.RequestObject) *webServerObject.ResponseObject {
	responseObj := webServerObject.NewResponseObject()

	//如果不是这几个方法，则要检测用户数据
	if requestObject.HTTPRequest.RequestURI != "/API/SysUser/Login" && requestObject.HTTPRequest.RequestURI != "/API/SysUser/Register" &&
		requestObject.HTTPRequest.RequestURI != "/API/SysUser/Identify" && requestObject.HTTPRequest.RequestURI != "/API/SysUser/Retrieve" {
		//根据用户名字判断过期时间
		userName, err := requestObject.GetStringVal("UserName")
		key, err2 := requestObject.GetStringVal("Token")
		if err != nil || err2 != nil {
			responseObj.SetResultStatus(webServerObject.DataError)
			return responseObj
		}

		//先检测key
		keyByte, err := EncrpytTool.RsaEncrypt([]byte(key))
		if err != nil {
			responseObj.SetResultStatus(webServerObject.SignError)
			return responseObj
		}

		//检测用户名是否是登录的
		stringArray := strings.Split(string(keyByte), ",")
		if stringArray[0] != userName {
			responseObj.SetResultStatus(webServerObject.SignError)
			return responseObj
		}

		//检测时间是否是用户上次登录时间
		timeStamp, err3 := strconv.ParseInt(stringArray[1], 10, 64)
		if err3 != nil {
			responseObj.SetResultStatus(webServerObject.SignError)
			return responseObj
		}

		if !CheckLoginTime(userName, timeStamp) {
			responseObj.SetResultStatus(webServerObject.SignError)
			return responseObj
		}

		//如果过期，返回过期提示
		if CheckPwdExpiredTime(userName) {
			responseObj.SetResultStatus(webServerObject.LoginIsOverTime)
			return responseObj
		} else {
			//如果没过期，返回新的过期时间
			UpdatePwdExpiredTime(userName)
			sysUserObj := GetItemByUserNameOrEmail(userName)
			if sysUserObj != nil {
				responseObj.AttachData["PwdExpiredTime"] = sysUserObj.PwdExpiredTime.UnixNano() / 1e6
			}
		}
	}

	return responseObj
}
