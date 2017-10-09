package model

import (
	"time"
)

// SysUserEmail 用户邮箱
type SysUserEmail struct {
	// Email 邮箱
	Email string

	// IdentifyCode 验证码
	IdentifyCode string

	// CRTime 更新时间
	CRTime time.Time
}

// NewSysUserEmail 新建用户邮箱
func NewSysUserEmail(email string, identifyCode string, cRTime time.Time) *SysUserEmail {
	return &SysUserEmail{
		Email:        email,
		IdentifyCode: identifyCode,
		CRTime:       cRTime,
	}
}
