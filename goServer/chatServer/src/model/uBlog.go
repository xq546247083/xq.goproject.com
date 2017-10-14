package model

import (
	"time"
)

// UBlog 用户博客表
type UBlog struct {
	// 主键
	ID string `gorm:"column:ID;primary_key"`

	// 用户id
	UserId string `gorm:"column:UserId"`

	// 标题
	Title string `gorm:"column:Title"`

	// 内容
	Content string `gorm:"column:Content"`

	// 标签（用，号隔开）
	Tag string `gorm:"column:Tag"`

	// @的用户
	ATUsers string `gorm:"column:ATUsers"`

	// 博客类型
	BlogType int32 `gorm:"column:BlogType"`

	// 状态【0：草稿，1：正常，2：删除，3：彻底删除】
	Status byte `gorm:"column:Status"`

	// 创建时间
	CrDate time.Time `gorm:"column:CrDate"`

	// 更新时间
	ReDate time.Time `gorm:"column:ReDate"`
}

//TableName 用户博客表
func (thisObj *UBlog) TableName() string {
	return "u_blog"
}
