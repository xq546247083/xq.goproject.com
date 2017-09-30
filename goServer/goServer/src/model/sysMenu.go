package model

// SysMenu 菜单表
type SysMenu struct {
	// 菜单标识
	MenuID int32 `gorm:"column:MenuID;primary_key"`

	// 上级ID
	ParentMenuID int32 `gorm:"column:ParentMenuID"`

	// 菜单名称
	MenuName string `gorm:"column:MenuName"`

	// 菜单地址
	MenuUrl string `gorm:"column:MenuUrl"`

	// 排序号
	SortOrder int32 `gorm:"column:SortOrder"`

	// 菜单图标路径（未用到）
	MenuIcon string `gorm:"column:MenuIcon"`

	// 常用菜单图标（未用到）
	BigMenuIcon string `gorm:"column:BigMenuIcon"`

	// 快捷键（未用到）
	ShortCut string `gorm:"column:ShortCut"`

	// 是否显示
	IsShow bool `gorm:"column:IsShow"`
}

//TableName 菜单表
func (thisObj *SysMenu) TableName() string {
	return "sys_menu"
}
