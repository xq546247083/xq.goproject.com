package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileBasedPipe()
	// inMemorySyncPipe()
}

func fileBasedPipe() {
	// 建立管道
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("Error: Couldn't create the named pipe: %s\n", err)
	}
	go func() {
		// 循环读取数据
		for {
			output := make([]byte, 100)
			n, err := reader.Read(output)
			if err != nil {
				fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
			}
			fmt.Printf("Read %d byte(s). [file-based pipe]\n", n)
		}
	}()

	// 给管道发送数据
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)

	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
	n, err = writer.Write(input)
	n, err = writer.Write(input)
	for {
	}
}

func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [in-memory pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}
