package main

import (
	"fmt"
	"os"
)

func main() {
	inByte := make([]byte, 16, 16)
	os.Stdin.Read(inByte)

	fmt.Println(string(inByte))
}
