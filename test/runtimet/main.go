package main

import (
	"time"
	"fmt"
	"runtime"

)

func main(){	
	go func(){
		runtime.Goexit()

		fmt.Println("xx")
	}()	

	fmt.Println(runtime.NumGoroutine())
	time.Sleep(2*time.Second)
	fmt.Println("x2x")
}