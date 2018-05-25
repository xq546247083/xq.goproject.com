package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err))
}
