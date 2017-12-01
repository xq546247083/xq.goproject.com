package mysqlSync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"xq.goproject.com/commonTools/logTool"
)

const (
	// 同步信息文件的名称
	con_InfoFile_Name = "syncInfo.txt"
)

// 同步信息项，保存已经处理过的文件的信息
type syncInfoItem struct {
	// 待处理文件的绝对路径
	FilePath string

	// 待处理文件的偏移量
	Offset int64
}

func newSyncInfoItem(filePath string, offset int64) *syncInfoItem {
	return &syncInfoItem{
		FilePath: filePath,
		Offset:   offset,
	}
}

// 同步信息对象
type syncInfoObject struct {
	// 同步文件
	file *os.File

	// 同步数据对象的唯一标识，用于进行重复判断
	identifier string

	// 同步信息项
	item *syncInfoItem
}

// 获取同步信息项
func (this *syncInfoObject) getSyncInfoItem() *syncInfoItem {
	return this.item
}

// 设置同步信息项
func (this *syncInfoObject) setSyncInfoItem(item *syncInfoItem) {
	this.item = item
}

// 更新同步信息项
func (this *syncInfoObject) syncData(item *syncInfoItem) {
	// 序列化数据
	data, err := json.Marshal(item)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "syncInfoObject.syncData.json.Marshal")
		err = fmt.Errorf("%s-Marshal syncInfoItem failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	// 覆盖写入
	this.file.Seek(0, 0)
	_, err = this.file.WriteString(string(data))
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "syncInfoObject.syncData.f.WriteString")
		err = fmt.Errorf("%s-Write data to syncInfo file failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	// 清理残留数据
	this.file.Truncate(int64(len(string(data))))

	// 更新同步信息项
	this.item = item
}

// 创建同步信息对象
// _dirPath:目录的路径
// _identifier:当前数据的唯一标识（可以使用数据库表名）
// 返回值:
// 同步信息对象
func newSyncInfoObject(dirPath, identifier string) *syncInfoObject {
	filePath := filepath.Join(dirPath, con_InfoFile_Name)

	// 打开Info文件, 如果没有就创建
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", identifier, "syncInfoObject.newSyncInfoObject.os.OpenFile")
		err = fmt.Errorf("%s-Open syncInfo file failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	// 读取文件内容
	var item *syncInfoItem
	content, err := ioutil.ReadAll(file)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", identifier, "syncInfoObject.newSyncInfoObject.fileUtil.ReadAll")
		err = fmt.Errorf("%s-Read syncInfo:%s failed:%s", prefix, filePath, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	// 初始化更新数据项
	if len(content) == 0 {
		item = newSyncInfoItem("", 0)
	} else {
		if err := json.Unmarshal([]byte(content), &item); err != nil {
			prefix := fmt.Sprintf("%s-%s", identifier, "syncInfoObject.newSyncInfoObject.json.Unmarshal")
			err = fmt.Errorf("%s-Unmarshal syncInfo data failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}
	}

	return &syncInfoObject{
		identifier: identifier,
		item:       item,
		file:       file,
	}
}
