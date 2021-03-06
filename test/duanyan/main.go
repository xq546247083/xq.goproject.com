package main

import "fmt"

// 测试断言以及字符串
func main() {
	var ti testInterface
	ti = new(chlidStruct)

	_, ok := ti.(*baseStruct)
	fmt.Println(ok)

	_, ok = ti.(*chlidStruct)
	fmt.Println(ok)
}

type testInterface interface {
	GetName()
}

type baseStruct struct {
}

func (this *baseStruct) GetName() {

}

type chlidStruct struct {
	baseStruct
}

func testFunc(temp *baseStruct) {

}

func testBase() {
	// 转换，强制转换
	var uintTemp uint32
	uintTemp = 1
	fmt.Println(int32(uintTemp))

	// 原生字符串
	numbers := `你好\r\n
1啊`
	// 解释性字符串
	numbers2 := "xx你好\r\n1xxx啊"

	// 数组,...指{}内有多少元素，就有多长
	numbers3 := [...]string{"xx你好1xxx啊", "x"}

	fmt.Println(numbers[0:6])

	fmt.Println(string(numbers))
	fmt.Println(string(numbers2))
	fmt.Println(numbers3)

	r1, r2 := interface{}(uintTemp).(int32)
	fmt.Println(r1, r2)

	// 断言，不同类型断言会失败
	fmt.Println(interface{}(uintTemp).(uint32))
	fmt.Println(interface{}(uintTemp).(int32))

	var intputStr string
	fmt.Scanln(intputStr)
}
