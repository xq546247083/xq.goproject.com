package model

// Novel 小说表
type Novel struct {
	// 小说名
	Name string `gorm:"column:Name;primary_key"`

	// 标题名
	Title string `gorm:"column:Title;primary_key"`

	// 源网站
	Source string `gorm:"column:Source;primary_key"`

	// 内容
	Content string `gorm:"column:Content"`
}

//TableName 小说表
func (thisObj *Novel) TableName() string {
	return "novel"
}

// NewNovel 新建小说
func NewNovel(name, title, source, content string) *Novel {
	return &Novel{
		Name:    name,
		Title:   title,
		Source:  source,
		Content: content,
	}
}
