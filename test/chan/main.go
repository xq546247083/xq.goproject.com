package main

import (
	"time"
	"fmt"
)

type intChan chan int

func main() {
	// 带1个缓存的通道
	temp:=make(intChan,1)
	readChan3(temp)
	
	temp<-0
	temp<-1
    // 关闭通道，会给当前通道的所有接受方一直发送关闭信息
	close(temp)

	time.Sleep(2*time.Second)
	go testCloseNil()

	for{

	}
}

// 用range来读取通道，简单版本
func readChan(intChan <-chan int){
	go func(){
		for i:=range intChan{
			fmt.Println(i)
		}

		fmt.Println("chan read close")
	}()
}

// 用状态来判断通道是否关闭。（推荐）
func readChan2(intChan <-chan int){
	go func(){
		for{
			i,ok:=<-intChan

			fmt.Println(i,ok)
			// 如果不break，会一直接受到（0,false）的数据
			if !ok{
				break
			}
		}

		fmt.Println("chan read close")
	}()
}

// 读取通道3。（select chan方式）
func readChan3(intChan <-chan int){
	go func(){
		// 读取状态
		readStatus:=false

		for{
			select {
			case i,ok := <-intChan:
				fmt.Println(i,ok)
				if !ok{
					readStatus=true
				}
			default:
				fmt.Println("sleep")
				time.Sleep(time.Millisecond*100)
			}

			// 如果读取完了，break
			if(readStatus){
				break
			}
		}

		fmt.Println("chan read close")
	}()
}

// 尝试关闭空通道
func testCloseNil(){
	// 捕获异常
	defer func(){
		if err:=recover();err!=nil{
			// 关闭一个nil的通道会抛出错误
			fmt.Println(err);
		}
	}()

	// 阻塞式通道
	temp2:=make(chan int)
	temp2=nil

    // 给一个nil的通道传入数据，会导致阻塞（这里将不会执行关闭通道的代码）
	temp2<-1
	close(temp2)
}