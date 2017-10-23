package fileTool

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var (
	mutex sync.Mutex
)

//GetCurrentPath 获取当前路径
// 返回值：
// 当前路径
func GetCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	fileAbsPath, _ := filepath.Abs(file)

	return filepath.Dir(fileAbsPath)
}

//GetFileList 获取目标文件列表（完整路径）
// path：文件夹路径
// 返回值：文件列表（完整路径）
func GetFileList(path string) ([]string, error) {
	files := make([]string, 0, 100)

	//遍历目录，获取所有文件列表
	filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//忽略目录
		if fi.IsDir() {
			return nil
		}

		// 添加到列表
		files = append(files, filename)

		return nil
	})

	return files, nil
}

// GetFileInfoList 获取目标文件列表（完整路径）
// path：文件夹路径
// 返回值：文件列表（完整路径）
func GetFileInfoList(path string) ([]os.FileInfo, error) {
	files := make([]os.FileInfo, 0, 100)

	//遍历目录，获取所有文件列表
	filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//忽略目录
		if fi.IsDir() {
			return nil
		}

		// 添加到列表
		files = append(files, fi)

		return nil
	})

	return files, nil
}

// 文件夹是否存在
func IsDirectoryExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

// 文件是否存在
func IsFileExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file.IsDir() == false, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

//ReadFileLineByLine 按行读取每一个文件的内容
// filename:文件的绝对路径
// 返回值：
// 行内容列表
// 错误信息
func ReadFileLineByLine(filename string) ([]string, error) {
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//读取文件
	lineList := make([]string, 0, 100)
	buf := bufio.NewReader(file)
	for {
		//按行读取
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}

		//将byte[]转换为string，并添加到列表中
		lineList = append(lineList, string(line))
	}

	return lineList, nil
}

//ReadFileContent 读取文件内容（字符串）
// filename：文件的绝对路径
// 返回值：
// 文件内容
// 错误信息
func ReadFileContent(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//ReadFileBytes 读取文件内容（字符数组）
// filename：文件的绝对路径
// 返回值：
// 文件内容
// 错误信息
func ReadFileBytes(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)

	return bytes, err
}

//DeleteFile 删除文件
// filename：文件的绝对路径
// 返回值：
// 错误对象
func DeleteFile(filename string) error {
	return os.Remove(filename)
}
