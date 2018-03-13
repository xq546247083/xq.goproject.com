package main

import (
	"runtime/debug"
	"runtime"
	"fmt"
	"sync/atomic"
	"sync"
	
)

func main() {
    var count int32
    addFunc:=func() interface{}{
        return atomic.AddInt32(&count,1)
    }
    pool:= sync.Pool{New:addFunc}

    fmt.Println(pool.Get())

    pool.Put(10)
    pool.Put(12)
    pool.Put(13)
    fmt.Println(pool.Get())

    debug.SetGCPercent(100)
    runtime.GC()

    fmt.Println(pool.Get())
}


