package model

import (
	"time"
)

// HistoryPrivate 私聊聊天记录
type HistoryPrivate struct {
	// ID
	ID int32 `gorm:"column:ID;primary_key"`

	// 私聊消息的接收者Id
	SysUserID string `gorm:"ConfigValue:SysUserID;primary_key"`

	// 聊天消息
	Message string `gorm:"column:Message"`

	// 语音信息
	Voice string `gorm:"column:Voice"`

	// 源用户Id
	FromSysUserID string `gorm:"column:FromSysUserID"`

	// 创建时间
	Crtime time.Time `gorm:"column:Crtime"`
}

//TableName 私聊聊天记录
func (thisObj *HistoryPrivate) TableName() string {
	return "history_private"
}
