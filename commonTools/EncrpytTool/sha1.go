package EncrpytTool

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

//Sha1String 对字符串进行SHA1加密
// s:输入字符串
// ifUpper:输出是否大写
// 返回值：md5加密后的字符串
func Sha1String(s string) string {
	if len(s) == 0 {
		panic(errors.New("input string can't be empty"))
	}

	return Sha1Bytes([]byte(s))
}

// Sha1Bytes 对字符数组进行SHA1加密
// b:输入字符数组
// ifUpper:输出是否大写
// 返回值：md5加密后的字符串
func Sha1Bytes(b []byte) string {
	if len(b) == 0 {
		panic(errors.New("input []byte can't be empty"))
	}

	sha1Instance := sha1.New()
	sha1Instance.Write(b)
	result := sha1Instance.Sum([]byte(""))

	return string(fmt.Sprintf("%x", result))
}

// Sha1BytesByNil 对字符数组进行SHA1加密()
// b:输入字符数组
// ifUpper:输出是否大写
// 返回值：md5加密后的字符串
func Sha1BytesByNil(b string) []byte {
	sha1Instance := sha1.New()
	io.WriteString(sha1Instance, b)
	return sha1Instance.Sum(nil)
}
