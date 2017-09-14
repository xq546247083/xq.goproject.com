package byteTool

import (
	"bytes"
	"encoding/binary"
)

//ByteToInt 转型
func ByteToInt(b []byte, order binary.ByteOrder) int {
	byteBuffer := bytes.NewBuffer(b)
	var result int
	binary.Read(byteBuffer, order, &result)

	return result
}

//ByteToInt32 转型
func ByteToInt32(b []byte, order binary.ByteOrder) int32 {
	byteBuffer := bytes.NewBuffer(b)
	var result int32
	binary.Read(byteBuffer, order, &result)

	return result
}

//ByteToString 转型
func ByteToString(b []byte) string {
	return string(b)
}
