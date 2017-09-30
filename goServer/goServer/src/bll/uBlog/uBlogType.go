package uBlog

import (
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

var (
	uBlogTypeMap = make(map[int32]*model.UBlogType)
)

func init() {
	initTool.RegisterInitFunc(initUBlogTypeData, initTool.I_NeedInit)
}

// 初始化数据
func initUBlogTypeData() error {
	uBlogTypeList, err := dal.UBlogTypeDALObj.GetAllList()
	if err != nil {
		return err
	}

	for _, item := range uBlogTypeList {
		uBlogTypeMap[item.ID] = item
	}

	return nil
}
