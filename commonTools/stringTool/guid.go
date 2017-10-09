package stringTool

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"xq.goproject.com/commonTools/EncrpytTool"
)

// GetNewGUID 获取新的GUID字符串
// 返回值：
// 新的GUID字符串
func GetNewGUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return EncrpytTool.Md5String(base64.URLEncoding.EncodeToString(b))
}

// IsGUIDEmpty 判断GUID是否为空
// guid：GUID
// 返回值：
// 是否为空
func IsGUIDEmpty(guid string) bool {
	if guid == "" || guid == "00000000-0000-0000-0000-000000000000" {
		return true
	}

	return false
}
