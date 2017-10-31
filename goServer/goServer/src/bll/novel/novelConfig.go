package novel

import (
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

// GetNovelConfigAllList 获取小说配置列表
func GetNovelConfigAllList() []*model.NovelConfig {
	novelConfigList, err := dal.NovelConfigDALObj.GetAllList()
	if err != nil {
		return nil
	}

	return novelConfigList
}
