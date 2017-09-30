package model

// SysRole 角色表
type SysRole struct {
	// 主键
	RoleID int32 `gorm:"column:RoleID;primary_key"`

	// 角色名称
	RoleName string `gorm:"column:RoleName"`

	// 菜单id（用,隔开）
	MenuIDS string `gorm:"column:MenuIDS"`

	// 是否默认角色
	IsDefault bool `gorm:"column:IsDefault"`

	// 是否是超级管理员角色
	IsSupper bool `gorm:"column:IsSupper"`

	// 描述
	Notes string `gorm:"column:Notes"`
}

//TableName 角色表
func (thisObj *SysRole) TableName() string {
	return "sys_role"
}
