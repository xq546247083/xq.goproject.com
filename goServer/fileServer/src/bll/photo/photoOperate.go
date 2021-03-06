package photo

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"xq.goproject.com/commonTools/fileTool"
)

//保存文件
func saveFile(fileName string, fileReader multipart.File) error {
	//判断文件夹是否存在，如果不存在，则创建
	exists, err := fileTool.IsDirectoryExists(uploadPath)
	if err != nil {
		return err
	}

	if !exists {
		if err = os.Mkdir(uploadPath, os.ModePerm); err != nil {
			return err
		}
	}

	//文件全路径
	fileFullPath := uploadPath + fileName
	//判断文件夹是否存在,如果存在，则追加
	fileExists, errT := fileTool.IsFileExists(fileFullPath)
	if errT != nil {
		return err
	}

	//如果不存在，则创建
	if !fileExists {
		//创建文件写入流
		fileWriter, err := os.Create(fileFullPath)
		if err != nil {
			return err
		}
		defer fileWriter.Close()

		//复制文件
		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}
	} else {
		errC := appendToFile(fileFullPath, fileReader)
		if errC != nil {
			return errC
		}
	}

	return nil
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, file multipart.File) error {
	// 以追加的模式，打开文件
	openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return err
	}

	// 查找文件末尾的偏移量
	locationNum, err := openFile.Seek(0, os.SEEK_END)
	if err != nil {
		return err
	}

	//读取字节
	data, err := readMultiPartFile(file)
	if err != nil {
		return err
	}

	// 从末尾的偏移量开始写入内容
	len, err := openFile.WriteAt(data, locationNum)
	if err != nil {
		return err
	}
	_ = len
	defer openFile.Close()

	return nil
}

//读取分块文件，返回字节
func readMultiPartFile(file multipart.File) ([]byte, error) {
	//从分片文件中读取字节,最大为1mb的分片数据
	buf := bytes.NewBuffer(make([]byte, 0, 1024*1024))
	n, err := buf.ReadFrom(file)
	return buf.Bytes()[:n], err
}
