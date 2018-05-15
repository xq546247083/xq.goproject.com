package main

import (
	"fmt"
)

// 测试了切片的一些操作
func main() {
	// 初始化一个切片
	dasd := make([]int32, 4, 5)
	dasd[0] = 0
	dasd[1] = 1
	dasd[2] = 2
	dasd[3] = 3
	// 最后一个参数，预设cap的最后一位（3个参数）
	a1 := dasd[0:2:3]
	// append一个数，会改变原来的切片
	a1 = append(a1, 9)
	// 超出cap，新建切片，不会影响原来的切片
	a1 = append(a1, 10)

	fmt.Println(dasd)
	fmt.Println(a1)

	a := 10
	f := func() int { a = a * 2; return 5 }
	// f函数在初始化的时候就执行了
	x := []int{a, f()}

	for a := 0; a < 2; a++ {
		fmt.Println(x)
	}
}

// 演示方法作为参数传入
type Encipher func(plaintext string) []byte

// encipher 作为一个方法类型
// 返回一个func
func GetItem(encipher Encipher) func(string) string {
	return func(dd string) string {
		return fmt.Sprintf("%x", encipher(dd))
	}
}
