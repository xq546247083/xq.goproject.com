package main

import (
	"fmt"
	"time"
)


func main() {
    timer:=time.NewTimer(3*time.Second)
    fmt.Println(time.Now())
    ss:=<-timer.C
    fmt.Println(time.Now())
    fmt.Println(ss)

    timeOutChan()
}

// 超时通道
func timeOutChan(){
    var chanInt chan int
    select{
        case e:=<-chanInt:
            fmt.Println(e)
        case <-time.After(2*time.Second):
            fmt.Println("timeout")
    }
}