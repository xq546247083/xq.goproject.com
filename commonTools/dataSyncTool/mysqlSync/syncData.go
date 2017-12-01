package mysqlSync

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"moqikaka.com/goutil/intAndBytesUtil"
	"moqikaka.com/goutil/logUtil"
	"xq.goproject.com/commonTools/fileTool"
)

const (
	con_Default_FileName = "00000000"
	con_Suffix           = "data"
	con_Header_Length    = 4
)

var (
	// 字节的大小端顺序
	byterOrder = binary.LittleEndian
)

// 用于大文件对象为新创建的文件命名
func newFileName(prefix, path string) string {
	fullName := filepath.Base(path)
	curFileName := strings.Split(fullName, ".")[0]
	curFileName_int, err := strconv.Atoi(curFileName)
	if err != nil {
		err = fmt.Errorf("%s-Convert newFileName:%s to int failed:%s", prefix, curFileName, err)
		logUtil.ErrorLog(err.Error())
		panic(err)
	}

	newFileName_int := curFileName_int + 1
	newFileName := fmt.Sprintf("%08d", newFileName_int)

	// 加上文件后缀
	newFileName = fmt.Sprintf("%s.%s", newFileName, con_Suffix)

	return newFileName
}

// 获取文件夹下的文件列表
func getDataFileList(dirPath string) []string {
	// 获取当前目录中所有的数据文件列表
	fileList, err := fileTool.GetFileList2(dirPath, "", con_Suffix)
	if err != nil {
		if os.IsNotExist(err) {
		} else {
			err = fmt.Errorf("%s/*.%s-Get file list failed:%s", dirPath, con_Suffix, err)
			logUtil.ErrorLog(err.Error())
			panic(err)
		}
	}

	// 如果文件数量大于1，则进行排序，以便于后续处理
	if len(fileList) > 1 {
		sort.Strings(fileList)
	}

	return fileList
}

// 同步数据对象
type syncDataObject struct {
	// 存放同步数据的文件夹路径
	dirPath string

	// 同步数据对象的唯一标识，用于进行重复判断
	identifier string

	// 保存数据的大文件对象
	bigFileObj *fileTool.BigFile

	mutex sync.Mutex
}

// 将数据写入同步数据对象
func (this *syncDataObject) write(data string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 获得数据内容的长度
	dataLength := len(data)

	// 将长度转化为字节数组
	header := intAndBytesUtil.Int32ToBytes(int32(dataLength), byterOrder)

	// 将头部与内容组合在一起
	message := append(header, data...)

	// 写入数据
	err := this.bigFileObj.WriteMessage(message)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "syncDataObject.write.bigFileObj.WriteMessage")
		err = fmt.Errorf("%s-Write message to big file object failed:%s", prefix, err)
		logUtil.ErrorLog(err.Error())
		panic(err)
	}
}

// 获取大文件对象的文件绝对路径
func (this *syncDataObject) filefullName() string {
	return filepath.Join(this.dirPath, this.bigFileObj.FileName())
}

// 创建同步数据对象
// _dirPath:目录的路径
// _identifier:当前数据的唯一标识（可以使用数据库表名）
// _maxFileSize:每个大文件的最大写入值（单位：Byte）
// 返回值:
// 同步数据对象
func newSyncDataObject(dirPath, identifier, fileName string, maxFileSize int) *syncDataObject {
	this := &syncDataObject{
		dirPath:    dirPath,
		identifier: identifier,
	}

	// 初始化大文件对象
	if fileName == "" {
		fileName = con_Default_FileName
	}
	bigFileObj, err := fileTool.NewBigFileWithNewFileNameFunc2(dirPath, "", fileName, maxFileSize, newFileName)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "syncDataObject.newSyncDataObject.fileTool.NewBigFileWithNewFileNameFunc")
		err = fmt.Errorf("%s-Create big file object failed:%s", prefix, err)
		logUtil.ErrorLog(err.Error())
		panic(err)
	}
	this.bigFileObj = bigFileObj

	return this
}
