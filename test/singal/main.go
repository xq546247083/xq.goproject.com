package main

import (
	"syscall"
	"fmt"
	"os/signal"
	"os"
)

func main(){	
	// 构造一个信号通道
	sigChan:=make(chan os.Signal,1)
	// 注册通道
	signal.Notify(sigChan)

	// 关闭通道的方式
	// signal.Stop(sigChan)
	// close(sigChan)

	// 通道的2个方式(关闭通道后，这种方式for会退出)
	for message:=range sigChan{
		fmt.Println(message)

		if  message == os.Interrupt {
			sigChan<-syscall.SIGKILL			
		}
	}

	// 通道的2个方式(关闭通道后，会一直循环，message会nil)
	for{
		message:=<-sigChan

		fmt.Println(message)

		if  message== os.Interrupt {
			go func(){
				sigChan<-syscall.SIGKILL
			}()
		
			// os.Exit(0)
		}
	}
}