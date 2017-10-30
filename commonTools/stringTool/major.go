package stringTool

import (
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

var (
	// Separator 分隔符
	Separator = "----------------------------------------------------------------"
)

// GetNewLine 获取新的一行
func GetNewLine() string {
	switch os := runtime.GOOS; os {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

// ToUpper 转换为大写
func ToUpper(content string) string {
	return strings.ToUpper(content)
}

// ToLower 转换为小写
func ToLower(content string) string {
	return strings.ToLower(content)
}

// IsEmpty 检查一个字符串是否是空字符串
//  content:上下文字符串
//  bool:true：空字符串 false：非空字符串
//  返回值：
func IsEmpty(content string) bool {
	if len(content) <= 0 {
		return true
	}

	return strings.IndexFunc(content, func(item rune) bool {
		return unicode.IsSpace(item) == false
	}) < 0
}

// StringToInt 转型,转换失败，返回-1
func StringToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err == nil {
		return n
	}

	return -1
}

// StringToInt32 转型,转换失败，返回-1
func StringToInt32(str string) int32 {
	n, err := strconv.Atoi(str)
	if err == nil {
		return int32(n)
	}

	return -1
}

// SplitToInt32List 切割字符串为int32数组
func SplitToInt32List(str string) []int32 {
	stringArray := strings.Split(str, ",")

	result := make([]int32, 0, len(stringArray))
	for _, item := range stringArray {
		//如果转换成功，则添加，否则，忽略
		num, err := strconv.Atoi(item)
		if err == nil {
			result = append(result, int32(num))
		}
	}

	return result
}

// GetURLDomainName 获取地址的域名
func GetURLDomainName(url string) string {
	if IsURL(url) {
		strs := strings.Split(url, "/")
		if len(strs) >= 3 {
			return strs[0] + "//" + strs[2]
		}
	}

	return ""
}
