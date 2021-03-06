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

// GetChapterItems 获取小说列表
func (thisObj *novelDAL) GetNovelItems() (novelList []*model.Novel, err error) {
	if err = DB.Raw("SELECT Name FROM novel GROUP BY NAME").Scan(&novelList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetNovelItems", novelDALName))
		return
	}

	return
}

// GetChapterItems 获取章节列表
func (thisObj *novelDAL) GetChapterItems(name string) (novelList []*model.Novel, err error) {
	if err = DB.Raw("SELECT Name,Title FROM novel  WHERE NAME =? GROUP BY Title;", name).Scan(&novelList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetChapterItems", novelDALName))
		return
	}

	return
}

// GetItems 获取数据
func (thisObj *novelDAL) GetItems(name, title string) (novelList []*model.Novel, err error) {
	if err = DB.Where("name = ? and title =? ", name, title).Find(&novelList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetItems", novelDALName))
		return
	}

	return
}

// GetItem 获取数据
func (thisObj *novelDAL) GetItem(name, title, source string) (novel *model.Novel, err error) {
	if err = DB.Where("name = ? and title =? and source=? ", name, title, source).First(&novel).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetItem", novelDALName))
		return
	}

	return
}

// GetAllList 获取数据
func (thisObj *novelDAL) GetAllList() (novelList []*model.Novel, err error) {
	if err = DB.Raw("SELECT Name,Title FROM novel ;").Scan(&novelList).Error; err != nil {
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
