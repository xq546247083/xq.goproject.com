package main

import (
	"time"
	"fmt"
	"sync"
	
)

// 测试once以及锁，以及其异常
func main() {
    locktest() 

    fmt.Scanln()
}

// 运行一次
func onectest(){
    var locker sync.Once
    
    go func(){
        locker.Do(func(){
            fmt.Println("xxx")
        })
    }()

    locker.Do(func(){
        fmt.Println("xxx")
    })
}

// 给已解锁的锁解锁，触发异常
func unlocktest(){
    defer func(){
        if err:=recover();err!=nil{
            fmt.Println(err)
        }
    }()

    var locker sync.Mutex
    locker.Lock()
    locker.Unlock()
    locker.Unlock()
}

// 测试锁
func locktest(){
    var locker sync.Mutex
    locker.Lock()

    for i:=1;i<=3;i++{
        go func(j int){
            fmt.Println("lock:",j)
            locker.Lock()
            fmt.Println("unlock:",j)
        }(i)
    }

    time.Sleep(1*time.Second)
    locker.Unlock()
}