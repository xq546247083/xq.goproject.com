package photo

import (
	"strings"
	"time"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/fileTool"
	"xq.goproject.com/commonTools/initTool"
	"xq.goproject.com/commonTools/stringTool"
)

var (
	uploadPath = configTool.UploadPath

	//用户的照片路径
	//key:用户名
	//key：照片类型
	//value：照片列表
	photoNameMap = make(map[string]map[photoType][]fileInfo)
)

func init() {
	initTool.RegisterInitFunc(initFileData, initTool.I_NeedInit)
}

// initFileData 初始化文件数据
func initFileData() error {
	fileList, err := fileTool.GetFileInfoList(uploadPath)
	if err != nil {
		return err
	}

	for _, fileInfoVal := range fileList {
		fileName := fileInfoVal.Name()
		//如果是照片，则解析文件
		if stringTool.IsImg(fileName) {
			strList := strings.Split(fileName, "_")
			if len(strList) < 3 {
				continue
			}

			userName := strList[0]
			photoTypeTemp := photoType(strList[1])
			if _, exists := photoNameMap[userName]; !exists {
				photoNameMap[userName] = make(map[photoType][]fileInfo)
			}

			fileInfoObj := fileInfo{FileName: fileName, DirName: uploadPath, ModName: fileInfoVal.ModTime()}
			photoNameMap[userName][photoTypeTemp] = append(photoNameMap[userName][photoTypeTemp], fileInfoObj)
		}
	}

	return nil
}

//添加图片数据到内存
func addPhoto(userName string, photoTypeTemp photoType, photoName string, modTime time.Time) {
	if _, exists := photoNameMap[userName]; !exists {
		photoNameMap[userName] = make(map[photoType][]fileInfo)
	}

	userAblumPhotos := photoNameMap[userName][ablum]
	for _, userPhoto := range userAblumPhotos {
		if userPhoto.FileName == photoName {
			return
		}
	}

	fileInfoObj := fileInfo{FileName: photoName, DirName: uploadPath, ModName: modTime}
	photoNameMap[userName][ablum] = append(photoNameMap[userName][ablum], fileInfoObj)
}

// 组装数据返回
func assembleToClient(userName string) []interface{} {
	clientInfo := make([]interface{}, 0, 32)

	_, exists := photoNameMap[userName]
	if !exists {
		return clientInfo
	}

	userAblumPhotos, exists := photoNameMap[userName][ablum]
	if !exists {
		return clientInfo
	}

	for _, userAblumPhoto := range userAblumPhotos {
		clientInfo = append(clientInfo, userAblumPhoto)
	}

	return clientInfo
}
