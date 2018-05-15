package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

// 测试pool
func main() {
	var count int32
	addFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: addFunc}

	// 如果没有值，调用new方法创建一个值
	fmt.Println(pool.Get())

	// 传入了3个值
	pool.Put(10)
	pool.Put(12)
	pool.Put(13)
	// 按顺序取值
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())

	debug.SetGCPercent(100)
	runtime.GC()

	// 回收以后，重新调用new方法取值
	fmt.Println(pool.Get())
}
