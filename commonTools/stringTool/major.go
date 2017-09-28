package stringTool

import (
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

var (
	//Separator 分隔符
	Separator = "----------------------------------------------------------------"
)

//GetNewLine 获取新的一行
func GetNewLine() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

//IsEmpty 检查一个字符串是否是空字符串
// content:上下文字符串
// 返回值：
// bool:true：空字符串 false：非空字符串
func IsEmpty(content string) bool {
	if len(content) <= 0 {
		return true
	}

	return strings.IndexFunc(content, func(item rune) bool {
		return unicode.IsSpace(item) == false
	}) < 0
}

//StringToInt 转型,转换失败，返回-1
func StringToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err == nil {
		return n
	}

	return -1
}
