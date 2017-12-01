package mysqlSync

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"xq.goproject.com/commonTools/fileTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/commonTools/timeTool"
)

const (
	// 同步日志文件的名称
	con_LogFile_Name = "syncLog.txt"
)

// 同步日志项，保存已经处理过的文件的信息
type syncLogItem struct {
	// 同步完成时间
	Time string

	// 同步文件名
	FileName string
}

// 构造同步日志项
// fileName: 同步文件名
func newSyncLogItem(fileName string) *syncLogItem {
	timeFormat := timeTool.Format(time.Now(), "yyyy-MM-dd HH:mm:ss")

	return &syncLogItem{
		Time:     timeFormat,
		FileName: fileName,
	}
}

// 同步日志对象
type syncLogObject struct {
	// 同步文件
	file *os.File

	// 文件绝对路径
	filePath string

	// 同步数据对象的唯一标识，用于进行重复判断
	identifier string

	// 存活时间 (单位:hour)
	survivalTime int

	// mutex
	mutex sync.Mutex
}

// 保存同步日志项
// fileName: 同步文件名
func (this *syncLogObject) savaLog(fileName string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 构造同步对象
	item := newSyncLogItem(fileName)

	// 序列化数据
	data, err := json.Marshal(item)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "syncLogObject.savaLog.json.Marshal")
		err = fmt.Errorf("%s-Marshal syncInfoItem failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	this.open(os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	defer this.close()

	// 写入消息（在结尾处增加一个换行符\n）
	message := fmt.Sprintf("%s\n", data)
	_, err = this.file.WriteString(message)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "syncLogObject.savaLog.f.WriteString")
		err = fmt.Errorf("%s-Write data to syncLog file failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}
}

// 定时日志清理
func (this *syncLogObject) logClean() {
	clean := func() {
		this.mutex.Lock()
		defer this.mutex.Unlock()

		// 读取所有Log信息
		lineList, err := fileTool.ReadFileLineByLine(this.filePath)
		if err != nil && !os.IsNotExist(err) {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "syncLogObject.LogClean.ReadFileLineByLine")
			err = fmt.Errorf("%s-ReadFileLineByLine failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}
		if len(lineList) == 0 {
			return
		}

		// 遍历，查看时间是否过期
		deleteLine := make([]string, 0)
		for _, line := range lineList {
			item := new(syncLogItem)
			json.Unmarshal([]byte(line), &item)
			logTime, err := timeTool.ToDateTime(item.Time)
			if err != nil {
				prefix := fmt.Sprintf("%s-%s", this.identifier, "syncLogObject.LogClean.ToDateTime")
				err = fmt.Errorf("time:%s, %s-ToDateTime failed:%s", item.Time, prefix, err)
				logTool.LogError(err.Error())
				panic(err)
			}
			intervalHour := (time.Now().Unix() - logTime.Unix()) / (60 * 60)
			if int(intervalHour) < this.survivalTime {
				break
			}

			// 如果存活时间超过，删除历史数据
			err = fileTool.DeleteFile(item.FileName)
			if err != nil && !os.IsNotExist(err) {
				prefix := fmt.Sprintf("%s-%s", this.identifier, "syncLogObject.LogClean.DeleteFile")
				err = fmt.Errorf("file:%s, %s- DeleteFile failed:%s", item.FileName, prefix, err)
				logTool.LogError(err.Error())
				panic(err)
			}
			deleteLine = append(deleteLine, line)
		}

		// 重置Log文件
		this.open(os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm|os.ModeTemporary)
		defer this.close()
		deleteLen := len(deleteLine)
		lineList = lineList[deleteLen:]
		content := make([]byte, 0)
		for _, line := range lineList {
			message := fmt.Sprintf("%s\n", line)
			content = append(content, []byte(message)...)
		}
		_, err = this.file.Write(content)
		if err != nil {
			prefix := fmt.Sprintf("%s-%s", this.identifier, "syncLogObject.LogClean.WriteString")
			err = fmt.Errorf("%s-Write data to syncLog file failed:%s", prefix, err)
			logTool.LogError(err.Error())
			panic(err)
		}
	}

	// 每隔1 hour,清理一次
	for {
		clean()

		time.Sleep(time.Hour * 1)
	}
}

// 打开文件
func (this *syncLogObject) open(flag int, perm os.FileMode) {
	var err error
	this.file, err = os.OpenFile(this.filePath, flag, perm)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "errorFile.newErrorFile.os.OpenFile")
		err = fmt.Errorf("%s-Open File failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}
}

// 关闭文件
func (this *syncLogObject) close() {
	this.file.Close()
}

// 创建同步日志对象
// dirPath:目录的路径
// identifier:当前数据的唯一标识（可以使用数据库表名）
// survivalTime：日志存活时间
// 返回值:
// 同步信息对象
func newSyncLogObject(dirPath, identifier string, survivalTime int) *syncLogObject {
	filePath := filepath.Join(dirPath, con_LogFile_Name)

	obj := &syncLogObject{
		identifier:   identifier,
		survivalTime: survivalTime,
		filePath:     filePath,
	}

	// 启动新协程定时清理日志
	go obj.logClean()

	return obj
}
