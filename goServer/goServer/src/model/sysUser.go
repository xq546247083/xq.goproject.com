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

//TableName 玩家表
func (thisObj *SysUser) TableName() string {
	return "sys_user"
}
