package model

// UBlogType 用户博客类型表
type UBlogType struct {
	// 主键
	ID int32 `gorm:"column:ID;primary_key"`

	// 类型名
	Name string `gorm:"column:Name"`

	// 图标
	Icon string `gorm:"column:Icon"`

	// 是否展示
	IsPublic bool `gorm:"column:IsPublic"`
}

//TableName 用户博客类型表
func (thisObj *UBlogType) TableName() string {
	return "u_blog_type"
}
