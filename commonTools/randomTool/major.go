package randomTool

import (
	"math/rand"
	"time"
)

// GetRandomStr 获取随机数据
func GetRandomStr(length int32) string {
	result := ""
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for a := int32(1); a <= length; a++ {
		index := random.Intn(len(str))
		result += string(str[index])
	}

	return result
}
