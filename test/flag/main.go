package main

import (
	"flag"
	"fmt"
)

// flag -N dsada
func main() {
	// 获取启动参数
	var paramStr string
	flag.StringVar(&paramStr, "N", "", "输入名字")
	if paramStr == "" {
		flag.Usage()
		return
	}

	fmt.Println(paramStr)
	fmt.Scanln()
}
