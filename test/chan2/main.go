package main

import (
	"time"
	"fmt"
)
func f1(ch chan int) {
    ch <- 1
}
func f2(ch chan int) {
    time.Sleep(time.Second * 10)
    ch <- 1
}

func main() {
    var ch1 = make(chan int,1)
    var ch2 = make(chan int)
    f1(ch1)
    go f2(ch2)

    for{
        select {
            case <-ch1:
                fmt.Println("The first case is selected.")
            case <-ch2:
                fmt.Println("The second case is selected.")
            // default:
            //     fmt.Println("The default case is selected.")
        }
        fmt.Println("xxxxxxxxxxxx.")
    }
    
}