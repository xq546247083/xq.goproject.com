package model

import (
	"time"
)

// HistoryWorld 世界聊天记录
type HistoryWorld struct {
	// ID
	ID int32 `gorm:"column:ID;primary_key"`

	// 聊天消息
	Message string `gorm:"column:Message"`

	// 语音信息
	Voice string `gorm:"column:Voice"`

	// 源用户Id
	FromSysUserID string `gorm:"column:FromSysUserID"`

	// 创建时间
	Crtime time.Time `gorm:"column:Crtime"`
}

//TableName 世界聊天记录
func (thisObj *HistoryWorld) TableName() string {
	return "history_world"
}
