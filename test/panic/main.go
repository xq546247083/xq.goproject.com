package main

import (
	"fmt"

	"xq.goproject.com/test/panic/testPanic"
)

func main() {
	getData()
}

func getData() {
	// 可以跨包捕获错误！
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	testPanic.GetData()
}
