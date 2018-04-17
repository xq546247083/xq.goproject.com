package main

import (
	"fmt"
	"time"
)

func main() {
    testTimer()
}

// 测试定时器
func testTimer(){
    timer:=time.NewTimer(3*time.Second)
    fmt.Println(time.Now())

    // 定时到了，传入当前时间
    ss:=<-timer.C
    fmt.Println(time.Now())
    fmt.Println(ss)
}

// 测试超时通道
func timeOutChan(){
    var chanInt chan int
    select{
        case e:=<-chanInt:
            fmt.Println(e)
        case <-time.After(2*time.Second):
            fmt.Println("timeout")
    }
}