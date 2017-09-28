package intTool

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

//IntToByte 转型(无效，因为系统无法判断读取的字节数)
func IntToByte(n int, order binary.ByteOrder) []byte {
	returnBuffer := bytes.NewBuffer([]byte{})
	binary.Write(returnBuffer, order, n)

	return returnBuffer.Bytes()
}

//Int32ToByte 转型
func Int32ToByte(n int32, order binary.ByteOrder) []byte {
	returnBuffer := bytes.NewBuffer([]byte{})
	binary.Write(returnBuffer, order, n)

	return returnBuffer.Bytes()
}

//Int64ToByte 转型
func Int64ToByte(n int64, order binary.ByteOrder) []byte {
	returnBuffer := bytes.NewBuffer([]byte{})
	binary.Write(returnBuffer, order, n)

	return returnBuffer.Bytes()
}

//IntToString  转型
func IntToString(n int) string {
	return strconv.Itoa(n)
}

//Int32ToString  转型
func Int32ToString(n int) string {
	return strconv.Itoa(n)
}
