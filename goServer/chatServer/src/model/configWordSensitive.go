package model

// ConfigWordSensitive 敏感词配置表
type ConfigWordSensitive struct {
	// 单词
	Word string `gorm:"column:Word;primary_key"`
}

//TableName 敏感词配置表
func (thisObj *ConfigWordSensitive) TableName() string {
	return "config_word_sensitive"
}
