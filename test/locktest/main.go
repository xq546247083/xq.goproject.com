package main

import (
	"time"
	"fmt"
	"sync"
	
)


func main() {
    onectest()   
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

    for{

    }
}

// 测试解锁(异常捕获，不能捕获到由解锁导致的错误)
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
            fmt.Println("locked:",j)
        }(i)
    }

    time.Sleep(1*time.Second)
    locker.Unlock()

    for{

    }
}