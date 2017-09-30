package model

import (
	"time"
)

// SysUserEmail 用户邮箱
type SysUserEmail struct {
	// Email 邮箱
	Email string

	// CRTime 更新时间
	CRTime time.Time
}
