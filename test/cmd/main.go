package main

import (
	"io"
	"bytes"
	"os/exec"
	"fmt"
)

func main(){	
	// 创建cmd，命令
	cmd1 := exec.Command("go", "env")

	// 建立输出管道
	cmdout1,err:=cmd1.StdoutPipe()
	if err!=nil{
		fmt.Println(err)
	}

	if err:=cmd1.Start();err!=nil{
		fmt.Println(err)
	}

	// 读取输出管道的信息
	var outputBuf bytes.Buffer
	for{
		tempOutPut:=make([]byte,5)
		n,err:=cmdout1.Read(tempOutPut)
		if err!=nil{
			if err==io.EOF{
				break
			}else{
				fmt.Println(err)
				break
			}
		}

		if n>0{
			outputBuf.Write(tempOutPut[:n])
		}
	}

	fmt.Println(outputBuf.String());

	fmt.Scanln()
}