package uBlog

import (
	"strings"

	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/goServer/chatServer/src/dal"
	"xq.goproject.com/goServer/chatServer/src/model"
	"xq.goproject.com/goServer/goServerModel/src/consts"
	"xq.goproject.com/goServer/goServerModel/src/enum"
)

var (
	uBlogMap = make(map[string]*model.UBlog)
)

func init() {
	initTool.RegisterInitFunc(initUBlogData, initTool.I_NeedInit)
}

// 初始化数据
func initUBlogData() error {
	uBlogList, err := dal.UBlogDALObj.GetAllList()
	if err != nil {
		return err
	}

	for _, uBlog := range uBlogList {
		//缓存部分数据，防止内存爆照
		length := 20
		if len(uBlog.Content) < 20 {
			length = len(uBlog.Content)
		}

		uBlog.Content = uBlog.Content[:length]
		uBlogMap[uBlog.ID] = uBlog
	}

	return nil
}

// 通过用户名获取博客数据
func getBlogByUser(sysUser *model.SysUser) []*model.UBlog {
	result := make([]*model.UBlog, 0, 32)
	for _, uBlog := range uBlogMap {
		if uBlog.UserId == sysUser.UserID {
			result = append(result, uBlog)
		}
	}

	return result
}

// 组装数据返回
func assembleToClient(sysUser *model.SysUser, blogTypeID int32, status int32, tagInfo string) []map[string]interface{} {
	clientInfoList := make([]map[string]interface{}, 0, 32)

	// 获取玩家所有博客
	blogs := getBlogByUser(sysUser)

	// 临时数据
	tempList := make([]*model.UBlog, 0, 32)

	// 如果没有这个类型的博客，那么则用状态筛选
	_, exists := uBlogTypeMap[blogTypeID]
	if exists {
		for _, uBlog := range blogs {
			if uBlog.BlogType == blogTypeID && uBlog.Status == byte(enum.BlogCommon) {
				tempList = append(tempList, uBlog)
			}
		}
	} else {
		for _, uBlog := range blogs {
			if uBlog.Status == byte(status) {
				tempList = append(tempList, uBlog)
			}
		}
	}

	// 返回的列表
	resultList := make([]*model.UBlog, 0, 32)

	// 通过标签筛选数据
	tagList := strings.Split(tagInfo, ",")
	if len(tagList) > 0 {
		for _, uBlog := range tempList {
			// 获取博客的tagList
			uBlogTagList := strings.Split(uBlog.Tag, ",")

			// 筛选
			for _, tag := range tagList {
				breakFlag := false

				for _, uBlogTag := range uBlogTagList {
					if tag == uBlogTag {
						resultList = append(resultList, uBlog)
						breakFlag = true
						break
					}
				}

				if breakFlag {
					break
				}
			}
		}
	} else {
		resultList = append(resultList, tempList...)
	}

	//组装数据
	for _, uBlog := range resultList {
		clientInfo := make(map[string]interface{})

		clientInfo[consts.ID] = uBlog.ID
		clientInfo[consts.Title] = uBlog.Title
		clientInfo[consts.Content] = uBlog.Content
		clientInfo[consts.Tag] = uBlog.Tag
		clientInfo[consts.ATUsers] = uBlog.ATUsers
		clientInfo[consts.BlogType] = uBlog.BlogType
		clientInfo[consts.Status] = uBlog.Status
		clientInfo[consts.CrDate] = uBlog.CrDate
		clientInfo[consts.ReDate] = uBlog.ReDate

		clientInfoList = append(clientInfoList, clientInfo)
	}

	return clientInfoList
}
