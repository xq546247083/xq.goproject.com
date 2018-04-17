package main

import (
	"fmt"
	"os"
)

// 获取线程id
func main(){
	pid := os.Getpid()
	fmt.Println(pid)

	ppid := os.Getppid()
	fmt.Println(ppid)
}