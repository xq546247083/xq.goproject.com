package novel

import (
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

// GetItems 获取小说章节列表
func GetItems(name, title string) []*model.Novel {
	novelList, err := dal.NovelDALObj.GetItems(name, title)
	if err != nil {
		return nil
	}

	return novelList
}

// IsExisItems 是否存在该小说章节
func IsExisItems(name, title string) bool {
	novels := GetItems(name, title)
	if novels == nil {
		return false
	}

	return len(novels) > 0
}

// GetItem 获取小说章节
func GetItem(name, title, source string) *model.Novel {
	novel, err := dal.NovelDALObj.GetItem(name, title, source)
	if err != nil {
		return nil
	}

	return novel
}

// IsExisItem 是否存在该小说章节
func IsExisItem(name, title, source string) bool {
	novels := GetItem(name, title, source)
	if novels == nil {
		return false
	}

	return true
}
