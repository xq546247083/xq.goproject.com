package main

import (
	"fmt"
	"io/ioutil"
	"os"

	cErrors "github.com/chai2010/errors"
)

func main() {
	wrapError()
}

func wrapError() {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		for i, x := range (cErrors.Wrap(err, "read failed")).(cErrors.Error).Caller() {
			fmt.Printf("caller:%d: %s\n", i, x.FuncName)
		}
	}
}

func judgeError() {
	_, err := os.Open("/no/such/file")

	// 判断错误类型
	fmt.Println(os.IsNotExist(err))
}
