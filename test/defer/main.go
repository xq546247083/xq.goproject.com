package main

import (
	"fmt"
)

func main() {
	fmt.Println(deferReturn2())
	return
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func deferReturn2() int {
	result := 0

	defer func() {
		result = 1
	}()

	defer func() {
		result = 2
	}()

	return result
}

func deferReturn() (result int) {
	defer func() {
		result = 1
	}()

	defer func() {
		result = 2
	}()

	result = 0

	return result
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// defer 是后进先出。
// panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
