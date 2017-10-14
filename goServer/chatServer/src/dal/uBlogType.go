package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/chatServer/src/model"
)

// uBlogTypeDAL 博客类型dal
type uBlogTypeDAL struct{}

var (
	// UBlogTypeDALObj dal数据对象
	UBlogTypeDALObj = new(uBlogTypeDAL)

	// DALName 连接对象名
	uBlogTypeDALName = "UBlogTypeDALObj"
)

// GetAllList 获取数据
func (thisObj *uBlogTypeDAL) GetAllList() (uBlogTypeList []*model.UBlogType, err error) {
	if err = DB.Find(&uBlogTypeList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", uBlogTypeDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *uBlogTypeDAL) SaveInfo(uBlogType *model.UBlogType, tempDB *gorm.DB) (err error) {
	if tempDB == nil {
		tempDB = DB
	}

	if err = tempDB.Save(uBlogType).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", uBlogTypeDALName))
		return
	}

	return nil
}
