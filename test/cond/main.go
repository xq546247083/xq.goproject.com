package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 条件变量测试案例
func main() {
	runtime.GOMAXPROCS(4)

	testCond()
}

func testCond() {
	c := sync.NewCond(&sync.Mutex{})
	condition := false

	// 通知方法
	go func() {
		time.Sleep(time.Second * 1)
		c.L.Lock()
		defer c.L.Unlock()

		fmt.Println("[1] 变更condition状态,并发出变更通知.")
		condition = true
		c.Signal() //c.Broadcast()
		fmt.Println("[1] 继续后续处理.")
	}()

	// 等待消息方法
	func(){
		// 锁住
		c.L.Lock()
		defer c.L.Unlock()
		fmt.Println("[2] condition..........1")

		// 如果条件为false，进入
		for !condition {
			fmt.Println("[2] condition..........2")
			// 在此处等待，并释放锁
			// 一旦接受到通知，则继续放行
			c.Wait()
			fmt.Println("[2] condition..........3")
		}
		fmt.Println("[2] condition..........4")
	}()

	fmt.Println("main end...")
}