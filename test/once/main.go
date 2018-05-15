package main

import (
	"fmt"
	"sync"
)

// 测试once
func main() {
	onectest()

	fmt.Scanln()
}

// 运行一次
func onectest() {
	var locker sync.Once

	go func() {
		locker.Do(func() {
			fmt.Println("xxx")
		})
	}()

	locker.Do(func() {
		fmt.Println("xxx")
	})
}
