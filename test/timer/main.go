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
    
}