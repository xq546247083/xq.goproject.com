package EncrpytTool

import "encoding/base64"

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
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
