package model

// ConfigWordForbid 禁止配置表
type ConfigWordForbid struct {
	// Word
	Word string `gorm:"column:Word;primary_key"`
}

//TableName 禁止配置表
func (thisObj *ConfigWordForbid) TableName() string {
	return "config_word_forbid"
}
