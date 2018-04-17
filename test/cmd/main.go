package main

import (
	"io"
	"bytes"
	"os/exec"
	"fmt"
)

// 测试输入cmd命令，输出对应的结果

// 管道是基于os.Pipe()
// os.Pipe()
func main(){
	// 创建cmd，命令
	cmd1 := exec.Command("go", "env")

	// 建立输出管道
	cmdout,err:=cmd1.StdoutPipe()
	if err!=nil{
		fmt.Println(err)
	}

	// 开始命令
	if err:=cmd1.Start();err!=nil{
		fmt.Println(err)
	}

	readCmd(cmdout)
	fmt.Scanln()
}

// 读取输出管道的信息
func readCmd(cmdOut io.ReadCloser){
	var outputBuf bytes.Buffer
	for{
		tempOutPut:=make([]byte,5)
		n,err:=cmdOut.Read(tempOutPut)
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
}