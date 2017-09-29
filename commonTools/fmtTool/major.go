package fmtTool

import (
	"fmt"

	"xq.goproject.com/commonTools/stringTool"
)

// Println 打印消息
func Println(msgs ...string) {
	printStr := ""
	for _, item := range msgs {
		printStr += stringTool.GetNewLine() + item
	}

	fmt.Println(printStr)
}
