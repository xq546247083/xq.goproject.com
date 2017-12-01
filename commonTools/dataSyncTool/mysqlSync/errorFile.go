package mysqlSync

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"xq.goproject.com/commonTools/fileTool"
	"xq.goproject.com/commonTools/logTool"
)

var (
	// 记录错误sql命令的文件名
	con_Error_FileName = "errorFile.txt"
)

// 定义处理错误命令的文件对象
type errorFile struct {
	// 错误文件
	file *os.File

	// 文件路径
	filePath string

	// 同步数据对象的唯一标识，用于进行重复判断
	identifier string
}

// 保存命令到错误文件
// command: sql命令
func (this *errorFile) saveCommand(command string) {
	this.open()
	defer this.close()

	// 覆盖写入
	this.file.Seek(0, 0)

	// 写入命令
	_, err := this.file.WriteString(command)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "errorFile.saveCommand")
		err = fmt.Errorf("%s-Write %s to file failed:%s", prefix, command, err)
		logTool.LogError(err.Error())
		panic(err)
	}

	// 清理残留数据
	this.file.Truncate(int64(len(command)))
}

// 读取文件中命令
func (this *errorFile) readCommand() string {
	this.open()
	defer this.close()

	this.file.Seek(0, 0)
	content, err := ioutil.ReadAll(this.file)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "errorFile.readCommand")
		err = fmt.Errorf("%s-Read command failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}
	return string(content)
}

// 打开文件
func (this *errorFile) open() {
	// 打开errorFile文件, 如果没有就创建
	var err error
	this.file, err = os.OpenFile(this.filePath, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		prefix := fmt.Sprintf("%s-%s", this.identifier, "errorFile.newErrorFile.os.OpenFile")
		err = fmt.Errorf("%s-Open File failed:%s", prefix, err)
		logTool.LogError(err.Error())
		panic(err)
	}
}

// 关闭文件
func (this *errorFile) close() {
	this.file.Close()
}

// 删除文件
func (this *errorFile) delete() {
	fileTool.DeleteFile(this.filePath)
}

// 构造错误文件对象
// dirPath:文件路径
func newErrorFile(dirPath string, identifier string) *errorFile {
	filePath := filepath.Join(dirPath, con_Error_FileName)
	return &errorFile{
		filePath:   filePath,
		identifier: identifier,
	}
}
