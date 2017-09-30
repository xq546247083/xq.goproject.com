package dal

import (
	"fmt"

	"xq.goproject.com/goServer/goServer/src/model"
)

// uBlogDAL 博客dal
type uBlogDAL struct{}

var (
	// UBlogDALObj dal数据对象
	UBlogDALObj = new(uBlogDAL)

	// uBlogDALName 连接对象名
	uBlogDALName = "UBlogDALObj"
)

// GetAllList 获取数据
func (thisObj *uBlogDAL) GetAllList() (uBlogList []*model.UBlog, err error) {
	if err = DB.Find(&uBlogList).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.GetAllList", uBlogDALName))
		return
	}

	return
}

// SaveInfo 保存数据
func (thisObj *uBlogDAL) SaveInfo(uBlog *model.UBlog) (err error) {
	if err = DB.Save(uBlog).Error; err != nil {
		writeErrorLog(err, fmt.Sprintf("%s.SaveInfo", uBlogDALName))
		return
	}

	return nil
}
