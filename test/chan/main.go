package main

import (
	"time"
	"fmt"
)

type intChan chan int

func main() {
	temp:=make(intChan,1)
	_=temp

	temp2:=make(chan int)
	temp2=nil
	_=temp2

	go func(){
		// time.Sleep(2*time.Second)
		for{
			i,ok:=<-temp

			fmt.Println("xxx")
			fmt.Println(i,ok)
			if !ok{
				break
			}

			
		}

		// for i:=range temp{
		//  	fmt.Println(i)
		// }
	}()

	temp<-0
	temp<-1
	temp<-2
	close(temp)

	for{

	}
}
