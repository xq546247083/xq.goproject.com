/*
条件变量 Cond 例子

Author: xcl
Date: 2015-11-29
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)

	testCond()
}

func testCond() {
	c := sync.NewCond(&sync.Mutex{})
	condition := false

	go func() {
		time.Sleep(time.Second * 1)
		c.L.Lock()
		fmt.Println("[1] 变更condition状态,并发出变更通知.")
		condition = true
		c.Signal() //c.Broadcast()
		fmt.Println("[1] 继续后续处理.")
		c.L.Unlock()

	}()

	c.L.Lock()
	fmt.Println("[2] condition..........1")
	for !condition {
		fmt.Println("[2] condition..........2")
		//等待Cond消息通知
		c.Wait()
		fmt.Println("[2] condition..........3")
	}
	fmt.Println("[2] condition..........4")
	c.L.Unlock()

	fmt.Println("main end...")
}