package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

// 测试获取传入的信号
func main() {
	go func() {
		time.Sleep(3 * time.Second)
		sendSignal()
	}()

	handleSignal()
}

// 处理信号
func handleSignal() {
	// 构造一个信号通道
	sigChan := make(chan os.Signal, 1)
	// 注册通道
	signal.Notify(sigChan)

	// 关闭通道的方式
	// signal.Stop(sigChan)
	// close(sigChan)

	// 通道的2个方式(关闭通道后，这种方式for会退出)
	for message := range sigChan {
		fmt.Println(message)

		if message == os.Interrupt {
			sigChan <- syscall.SIGKILL
		}
	}

	// 通道的2个方式(关闭通道后，会一直循环，message会nil)
	for {
		message := <-sigChan

		fmt.Println(message)

		if message == os.Interrupt {
			go func() {
				sigChan <- syscall.SIGKILL
			}()

			// os.Exit(0)
		}
	}
}

// 发送信号
func sendSignal() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal Error: %s\n", err)
			debug.PrintStack()
		}
	}()

	// 找到自己的pid
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		fmt.Printf("Process Finding Error: %s\n", err)
		return
	}

	// 发送信号
	sig := syscall.SIGKILL
	err = proc.Signal(sig)
	if err != nil {
		fmt.Printf("Signal Sending Error: %s\n", err)
		return
	}
}
