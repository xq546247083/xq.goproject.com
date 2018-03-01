package main

import (
	"fmt"
	
)

func main(){

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
	numbers3 := [...]string{"xx你好1xxx啊","x"}

	fmt.Println(numbers[0:6])

	fmt.Println(string(numbers))
	fmt.Println(string(numbers2))
	fmt.Println(numbers3)

	r1,r2:=interface{}(uintTemp).(int32)
	fmt.Println(r1,r2)

	// 断言，不同类型断言会失败
	fmt.Println(interface{}(uintTemp).(uint32))
	fmt.Println(interface{}(uintTemp).(int32))

	var  intputStr string 
	fmt.Scanln(intputStr)
}