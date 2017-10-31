package model

// NovelConfig 小说配置表
type NovelConfig struct {
	// 网址名称
	SiteName string `gorm:"column:SiteName;primary_key"`

	// 小说名字
	NovelName string `gorm:"column:NovelName;primary_key"`

	// 小说地址
	NovelAddress string `gorm:"column:NovelAddress"`
}

//TableName 小说表
func (thisObj *NovelConfig) TableName() string {
	return "novel_config"
}

// NewNovelConfig 新建小说配置
func NewNovelConfig(siteName, novelName, novelAddress string) *NovelConfig {
	return &NovelConfig{
		SiteName:     siteName,
		NovelName:    novelName,
		NovelAddress: novelAddress,
	}
}
