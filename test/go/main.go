package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	test()
}

func test() {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

	for {
	}
}

func test1() {
	for i := 1; i <= 10; i++ {
		// wg等待
		wg.Add(2)

		// 初始化变量
		x := 0
		y := 0

		go func() {
			y = 1
			fmt.Print("x:", x, " ")
			wg.Done()
		}()
		// 执行操作
		go func() {
			x = 1
			fmt.Print("y:", y, " ")
			wg.Done()
		}()

		// 等待
		wg.Wait()
		fmt.Println("")
	}
}
