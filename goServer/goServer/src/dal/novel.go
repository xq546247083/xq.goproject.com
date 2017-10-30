package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/goServer/src/model"
)

// novelDAL 小说dal
type novelDAL struct{}

var (
	// NovelDALObj dal数据对象
	NovelDALObj = new(novelDAL)

	// novelDALName 连接对象名
	novelDALName = "NovelDALObj"
)

// GetAllList 获取数据
func (thisObj *novelDAL) GetAllList() (novelList []*model.Novel, err error) {
	if err = DB.Find(&novelList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", novelDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *novelDAL) SaveInfo(novel *model.Novel, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(novel).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", novelDALName))
		return
	}

	return nil
}
