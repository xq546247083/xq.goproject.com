package logTool

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/fileTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/commonTools/timeTool"
)

var (
	//LogPath 日志地址
	LogPath string

	//logInfoFlag 是否写info日志
	logInfoFlag = configTool.LogInfoFlag

	//logDebugFlag 是否写Debug日志
	logDebugFlag = configTool.LogDebugFlag

	//logWarnFlag 是否写Warn日志
	logWarnFlag = configTool.LogWarnFlag

	//logErrorFlag 是否写Error日志
	logErrorFlag = configTool.LogErrorFlag

	//logFatalFlag 是否写Fatal日志
	logFatalFlag = configTool.LogFatalFlag
)

//init 开始写日志
func init() {
	//获取路径
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	LogPath = filepath.Join(filepath.Dir(path), configTool.LogPath)
}

//writeLog 写日志
func writeLog(logType LogType, content string) {
	if LogPath == "" {
		LogPath = "Log"
	}

	//判定是否写日志
	if logType.ToString() == "Info" && logInfoFlag == false {
		return
	}
	if logType.ToString() == "Debug" && logDebugFlag == false {
		return
	}
	if logType.ToString() == "Warn" && logWarnFlag == false {
		return
	}
	if logType.ToString() == "Error" && logErrorFlag == false {
		return
	}
	if logType.ToString() == "Fatal" && logFatalFlag == false {
		return
	}
	// 获取当前时间
	now := time.Now()

	// 获得年、月、日、时的字符串形式
	yearString := strconv.Itoa(now.Year())
	monthString := strconv.Itoa(int(now.Month()))
	dayString := strconv.Itoa(now.Day())
	hourString := strconv.Itoa(now.Hour())

	// 构造文件路径和文件名
	filePath := filepath.Join(LogPath, yearString, monthString)

	// 得到fileName
	fileName := fmt.Sprintf("%s-%s-%s-%s-%s.txt", yearString, monthString, dayString, hourString, logType.ToString())
	fileName = filepath.Join(filePath, fileName)

	// 判断文件夹是否存在，如果不存在则创建
	if flag, err := fileTool.IsDirectoryExists(filePath); err == nil && !flag {
		if err := os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary); err != nil {
			return
		}
	} else if err != nil {
		return
	}

	// 打开文件(如果文件存在就以读写模式打开，并追加写入；如果文件不存在就创建，然后以写模式打开。)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 写入内容
	f.WriteString(content + "\r\n")
}

//LogObject 记录对象的日志
func LogObject(logType LogType, contentObj ...interface{}) {
	content := fmt.Sprintf("【%s】", timeTool.GetNowStr())
	for _, item := range contentObj {
		contentByte, _ := json.Marshal(item)
		content += stringTool.GetNewLine() + string(contentByte)
	}

	//如果是错误日志，获取堆栈信息
	if logType == Error {
		for skip := 1; skip <= 10; skip++ {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			content += stringTool.GetNewLine()
			content += fmt.Sprintf("skip = %d, file = %s, line = %d", skip, file, line)

		}
	}

	content += stringTool.GetNewLine() + stringTool.Separator

	writeLog(logType, content)
}

//Log 记录日志
func Log(logType LogType, contentStr ...string) {
	content := fmt.Sprintf("【%s】", timeTool.GetNowStr())

	for _, item := range contentStr {
		content += stringTool.GetNewLine() + item
	}

	//如果是错误日志，获取堆栈信息
	if logType == Error {
		for skip := 1; skip <= 10; skip++ {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			content += stringTool.GetNewLine()
			content += fmt.Sprintf("skip = %d, file = %s, line = %d", skip, file, line)
		}
	}

	content += stringTool.GetNewLine() + stringTool.Separator

	writeLog(logType, content)
}

//LogInfo 记录消息日志
func LogInfo(contentStr ...string) {
	Log(Info, contentStr...)
}

//LogDebug 记录调试日志
func LogDebug(contentStr ...string) {
	Log(Debug, contentStr...)
}

//LogWarn 记录警告日志
func LogWarn(contentStr ...string) {
	Log(Warn, contentStr...)
}

//LogError 记录错误日志
func LogError(contentStr ...string) {
	Log(Error, contentStr...)
}

//LogFatal 记录致命错误日志
func LogFatal(contentStr ...string) {
	Log(Fatal, contentStr...)
}

//LogObjectInfo 记录消息日志
func LogObjectInfo(objects ...interface{}) {
	LogObject(Info, objects...)
}

//LogObjectDebug 记录调试日志
func LogObjectDebug(objects ...interface{}) {
	LogObject(Debug, objects...)
}

//LogObjectWarn 记录警告日志
func LogObjectWarn(objects ...interface{}) {
	LogObject(Warn, objects...)
}

//LogObjectError 记录错误日志
func LogObjectError(objects ...interface{}) {
	LogObject(Error, objects...)
}

//LogObjectFatal 记录致命错误日志
func LogObjectFatal(objects ...interface{}) {
	LogObject(Fatal, objects...)
}

//记录未知错误日志（不需要的方式，暂时保留）
// r：recover对象
// 返回值：无
func logUnknownError(r interface{}, args ...string) {
	// 组装所有需要写入的内容
	logInfo := fmt.Sprintf("【%s】", timeTool.GetNowStr())

	contentByte, _ := json.Marshal(r)

	logInfo += stringTool.GetNewLine() + fmt.Sprintf("通过recover捕捉到的未处理异常：%s", string(contentByte))

	// 获取附加信息
	if len(args) > 0 {
		logInfo += stringTool.GetNewLine() + fmt.Sprintf("附加信息：%s", strings.Join(args, ";"))
	}

	// 获取堆栈信息
	for skip := 1; skip <= 10; skip++ {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		logInfo += stringTool.GetNewLine()
		logInfo += fmt.Sprintf("skip = %d, file = %s, line = %d", skip, file, line)
	}

	logInfo += stringTool.GetNewLine() + stringTool.Separator
	// 构造对象并添加到队列中
	writeLog(Error, logInfo)
}
