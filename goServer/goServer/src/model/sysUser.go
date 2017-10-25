package model

import (
	"time"
)

// SysUser 玩家表
type SysUser struct {
	// 主键
	UserID string `gorm:"column:UserID;primary_key"`

	// 登录ID
	UserName string `gorm:"column:UserName"`

	// 用户真实姓名
	FullName string `gorm:"column:FullName"`

	// 登陆密码
	Password string `gorm:"column:Password"`

	// 密码过期时间
	PwdExpiredTime time.Time `gorm:"column:PwdExpiredTime"`

	// 性别 1男0女
	Sex bool `gorm:"column:Sex"`

	// 工作电话
	Phone string `gorm:"column:Phone"`

	// 电子邮箱
	Email string `gorm:"column:Email"`

	// 头像
	HeadImgage string `gorm:"column:HeadImgage"`

	// 状态 1 启用 2禁用 3已删
	Status int32 `gorm:"column:Status"`

	// 登录次数
	LoginCount int32 `gorm:"column:LoginCount"`

	// 最后登录时间
	LastLoginTime time.Time `gorm:"column:LastLoginTime"`

	// 公司ID
	LastLoginIP string `gorm:"column:LastLoginIP"`

	// 角色ID（可以多个）
	RoleIDs string `gorm:"column:RoleIDs"`

	// 创建日期
	CreateTime time.Time `gorm:"column:CreateTime"`
}

// TableName 玩家表
func (thisObj *SysUser) TableName() string {
	return "sys_user"
}

// NewSysUser 新建用户
func NewSysUser(userID string, userName string, password string, fullName string, sex bool, phone string, email string, status int32,
	loginCount int32, roleIDs string, createTime time.Time, pwdExpiredTime time.Time) *SysUser {
	return &SysUser{
		UserID:         userID,
		UserName:       userName,
		Password:       password,
		FullName:       fullName,
		Sex:            sex,
		Phone:          phone,
		Email:          email,
		Status:         status,
		LoginCount:     loginCount,
		RoleIDs:        roleIDs,
		CreateTime:     createTime,
		PwdExpiredTime: pwdExpiredTime,
	}

}
