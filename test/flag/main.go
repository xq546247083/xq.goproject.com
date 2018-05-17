package main

import (
	"flag"
	"fmt"
)

// flag -N dsada
func main() {
	// 获取启动参数
	var paramStr string
	flag.StringVar(&paramStr, "N", "默认字符串", "输入名字")
	flag.Parse()

	fmt.Println(paramStr)
	fmt.Scanln()
}
