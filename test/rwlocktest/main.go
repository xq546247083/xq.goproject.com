package main

import (
	"time"
	"fmt"
	"sync"
	
)


func main() {
    locktest()   
}

// 测试锁
func locktest(){
    var locker sync.RWMutex
    locker.Lock()
    for i:=1;i<=3;i++{
        go func(j int){
            fmt.Println("lock:",j)
            locker.Lock()
            fmt.Println("locked:",j)
        }(i)
    }

    time.Sleep(1*time.Second)
    locker.Unlock()

    for{

    }
}