package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	str := strconv.Itoa(x)

	var resultStr = ""
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "-" {
			resultStr = string(str[i]) + resultStr
		} else {
			resultStr += string(str[i])
		}
	}

	result64, err := strconv.Atoi(resultStr)
	if err != nil {
		return 0
	}

	if result64 > 2147483647 || result64 < -2147483648 {
		return 0
	}

	return int(result64)
}
