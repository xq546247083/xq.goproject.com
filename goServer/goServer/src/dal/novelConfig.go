package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/goServer/src/model"
)

// novelConfigDAL 小说配置dal
type novelConfigDAL struct{}

var (
	// NovelConfigDALObj dal数据对象
	NovelConfigDALObj = new(novelConfigDAL)

	// novelConfigDALName 连接对象名
	novelConfigDALName = "NovelConfigDALObj"
)

// GetAllList 获取数据
func (thisObj *novelConfigDAL) GetAllList() (novelConfigList []*model.NovelConfig, err error) {
	if err = DB.Find(&novelConfigList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", novelConfigDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *novelConfigDAL) SaveInfo(novelConfig *model.NovelConfig, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(novelConfig).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", novelConfigDALName))
		return
	}

	return nil
}
