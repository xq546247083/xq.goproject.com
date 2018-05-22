package main

import (
	"fmt"
	"reflect"
)

func main() {
	iTemp := 1
	iInterface := interface{}(iTemp)

	// 判断类型
	switch iInterface.(type) {
	case int:
		fmt.Println(iInterface.(int))
	default:
		// 获取类型
		fmt.Println(reflect.TypeOf(iInterface), iInterface)
	}
}
