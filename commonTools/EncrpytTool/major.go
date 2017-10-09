package EncrpytTool

import (
	"fmt"
)

var (
	// DefaultPreKey 默认前置key
	DefaultPreKey = "!1@2#3$4%5^6"

	// DefaultProKey 默认后置key
	DefaultProKey = "9!8@7#6$5%3^"
)

// Encrypt 加密字符串
func Encrypt(content string) string {
	content = fmt.Sprintf("%s%s%s", DefaultPreKey, content, DefaultProKey)
	return Md5String(Sha1String(content))
}
