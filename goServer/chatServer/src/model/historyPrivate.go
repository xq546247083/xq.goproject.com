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

	// 私聊消息的接收者Name
	SysUserName string `gorm:"ConfigValue:SysUserName;"`

	// 聊天消息
	Message string `gorm:"column:Message"`

	// 语音信息
	Voice string `gorm:"column:Voice"`

	// 源用户Id
	FromSysUserID string `gorm:"column:FromSysUserID"`

	// 源用户Name
	FromSysUserName string `gorm:"column:FromSysUserName"`

	// 创建时间
	Crtime time.Time `gorm:"column:Crtime"`
}

//TableName 私聊聊天记录
func (thisObj *HistoryPrivate) TableName() string {
	return "history_private"
}

// NewHistoryPrivate 新建私聊聊天记录
func NewHistoryPrivate(sysUserID, sysUserName, message, voice, fromSysUserID, fromSysUserName string) *HistoryPrivate {
	return &HistoryPrivate{
		SysUserID:       sysUserID,
		SysUserName:     sysUserName,
		Message:         message,
		Voice:           voice,
		FromSysUserID:   fromSysUserID,
		FromSysUserName: fromSysUserName,
		Crtime:          time.Now(),
	}
}
