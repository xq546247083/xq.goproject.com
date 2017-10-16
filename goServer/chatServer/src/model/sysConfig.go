package model

// SysConfig 系统配置
type SysConfig struct {
	// 配置Key
	ConfigKey int32 `gorm:"column:ConfigKey;primary_key"`

	// 配置值
	ConfigValue string `gorm:"ConfigValue:ConfigValue"`

	// 配置描述
	ConfigDesc string `gorm:"column:ConfigDesc"`
}

//TableName 系统配置
func (thisObj *SysConfig) TableName() string {
	return "sys_config"
}
