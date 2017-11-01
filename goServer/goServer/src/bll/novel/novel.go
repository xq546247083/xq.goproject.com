package novel

import (
	"sort"
	"strconv"
	"strings"

	"xq.goproject.com/commonTools/logTool"

	"xq.goproject.com/goServer/goServer/src/dal"
	"xq.goproject.com/goServer/goServer/src/model"
)

// GetNovelItems 获取小说列表
func GetNovelItems() []*model.Novel {
	novelList, err := dal.NovelDALObj.GetNovelItems()
	if err != nil {
		return nil
	}

	return novelList
}

// GetChapterItems 获取小说章节列表
func GetChapterItems(name string) []*model.Novel {
	novelList, err := dal.NovelDALObj.GetChapterItems(name)
	if err != nil {
		return nil
	}

	return novelList
}

// GetNovelInfos 获取小说信息
func GetNovelInfos(name, title string, flag int32) []*model.Novel {
	novelList, err := dal.NovelDALObj.GetChapterItems(name)
	if err != nil {
		return nil
	}

	//按照顺序升序排序
	sort.Slice(novelList, func(i, j int) bool {
		before := strings.Split(novelList[i].Title, ".")[0]
		after := strings.Split(novelList[j].Title, ".")[0]

		beforeInt, err := strconv.Atoi(before)
		afterInt, err2 := strconv.Atoi(after)
		if err != nil || err2 != nil {
			return false
		}

		return beforeInt < afterInt
	})

	curenNum := 0
	for index, novel := range novelList {
		if novel.Title == title {
			curenNum = index
		}
	}

	logTool.LogInfo(strconv.Itoa(curenNum))
	if int32(curenNum)+flag < 0 {
		return nil
	} else if curenNum+int(flag) >= len(novelList) {
		return nil
	}

	returnList, err2 := dal.NovelDALObj.GetItems(name, novelList[int32(curenNum)+flag].Title)
	if err2 != nil {
		return nil
	}

	return returnList
}

// GetItems 获取小说章节列表
func GetItems(name, title string) []*model.Novel {
	novelList, err := dal.NovelDALObj.GetItems(name, title)
	if err != nil {
		return nil
	}

	return novelList
}

// GetItem 获取小说章节
func GetItem(name, title, source string) *model.Novel {
	novel, err := dal.NovelDALObj.GetItem(name, title, source)
	if err != nil {
		return nil
	}

	return novel
}

// IsExisItems 是否存在该小说章节
func IsExisItems(name, title string) bool {
	novels := GetItems(name, title)
	if novels == nil {
		return false
	}

	return len(novels) > 0
}

// IsExisItem 是否存在该小说章节
func IsExisItem(name, title, source string) bool {
	novels := GetItem(name, title, source)
	if novels == nil {
		return false
	}

	return true
}

// 组装数据返回
func assembleNovelListToClient() []string {
	novelList := GetNovelItems()
	clientInfo := make([]string, 0, len(novelList))
	for _, novel := range novelList {
		clientInfo = append(clientInfo, novel.Name)
	}

	return clientInfo
}

// 组装章节数据返回
func assembleChapterListToClient(name string, num int32) []string {
	chapterList := GetChapterItems(name)
	clientInfo := make([]string, 0, len(chapterList))

	//循环章节，如果不存在，这添加章节
	for _, chapter := range chapterList {
		flag := false
		for _, clientSingleInfo := range clientInfo {
			if clientSingleInfo == chapter.Title {
				flag = true
			}
		}

		if !flag {
			clientInfo = append(clientInfo, chapter.Title)
		}
	}

	//按照顺序升序排序
	sort.Slice(clientInfo, func(i, j int) bool {
		before := strings.Split(clientInfo[i], ".")[0]
		after := strings.Split(clientInfo[j], ".")[0]

		beforeInt, err := strconv.Atoi(before)
		afterInt, err2 := strconv.Atoi(after)
		if err != nil || err2 != nil {
			return false
		}

		return beforeInt > afterInt
	})

	aLength := (num - 1) * 1000
	if int(aLength) > len(clientInfo) {
		return clientInfo[:0]
	}

	bLength := num * 1000
	if int(bLength) > len(clientInfo) {
		bLength = int32(len(clientInfo))
	}

	return clientInfo[aLength:bLength]
}
