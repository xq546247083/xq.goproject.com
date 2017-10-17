package EncrpytTool

import "encoding/base64"

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var coder = base64.NewEncoding(base64Table)

// Base64Encrypt 加密
func Base64Encrypt(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

// Base64Decrypt 解密
func Base64Decrypt(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

// Base64EncryptByStd 加密
func Base64EncryptByStd(src []byte) []byte {
	accept := make([]byte, 28)
	base64.StdEncoding.Encode(accept, src)

	return accept
}
