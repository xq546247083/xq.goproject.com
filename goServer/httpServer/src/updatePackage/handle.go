package updatePackage

import (
	"fmt"
	"os"
	"strings"
	"time"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/fileTool"
	"xq.goproject.com/commonTools/logTool"
)

var (
	ingoreErrorStr = "The process cannot access the file because it is being used by another process"
)

// 开启文件监控
func Monitor(monitorServerChan chan<- bool) {
	// 获取压缩包名
	updatePackageName := configTool.WebMainPath + "/" + configTool.UpdatePackage

	for {
		fileExiststFlag, err := fileTool.IsFileExists(updatePackageName)
		if err != nil {
			logTool.LogError(fmt.Sprintf("获取文件失败!err:%s", err))
			continue
		}

		if fileExiststFlag && TryOpen(updatePackageName) {
			// 解压缩期间关闭服务器
			logTool.LogInfo(fmt.Sprintf("开始解压缩升级包!"))
			monitorServerChan <- false
			err = fileTool.Untar(updatePackageName, "")

			if err != nil {
				if !strings.Contains(err.Error(), ingoreErrorStr) {
					logTool.LogError(fmt.Sprintf("解压缩失败!err:%s", err))
				}
			} else {
				fileTool.DeleteFile(updatePackageName)
				logTool.LogInfo(fmt.Sprintf("解压缩升级包成功!"))
			}

			monitorServerChan <- true
		}

		// 睡一秒
		time.Sleep(time.Second)
	}
}

// 尝试打开文件
func TryOpen(source string) bool {
	reader, err := os.Open(source)
	defer reader.Close()
	if err != nil {
		return false
	}

	return true
}
