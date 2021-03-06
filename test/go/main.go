package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	// go say("world")
	// say("hello")

	testSelect()
}

func testSelect() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func(chanInt chan<- int) {
		time.Sleep(500 * time.Second)
		chanInt <- 1
	}(chan1)

	select {
	case a := <-chan1:
		fmt.Println("a", a)
	case b := <-chan2:
		fmt.Println("b", b)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	fmt.Print(runtime.NumGoroutine())
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
