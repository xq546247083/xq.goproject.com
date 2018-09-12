package main

import "fmt"

func main() {
	fmt.Println(addBinary("0", "0"))
}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)

	// 获取最长的长度
	length := aLen
	if bLen > aLen {
		length = bLen
	}

	result := ""

	// 当前位为1的数量
	e1Num := 0
	for i := 1; i <= length; i++ {
		if aLen-i >= 0 {
			if string(a[aLen-i]) == "1" {
				e1Num++
			}
		}
		if bLen-i >= 0 {
			if string(b[bLen-i]) == "1" {
				e1Num++
			}
		}

		// 如果三位都是1，那么进位为1
		if e1Num == 3 {
			result = "1" + result
			e1Num = 1
		} else if e1Num == 2 {
			result = "0" + result
			e1Num = 1
		} else if e1Num == 1 {
			result = "1" + result
			e1Num = 0
		} else {
			result = "0" + result
			e1Num = 0
		}
	}

	if e1Num == 1 {
		result = "1" + result
	}

	return result
}
