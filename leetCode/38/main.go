package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		return say(countAndSay(n - 1))
	}
}

func say(origin string) string {
	result := ""

	iLen := 1
	for i := 0; i < len(origin); i++ {
		if i == len(origin)-1 {
			result += strconv.Itoa(iLen) + origin[i:i+1]
			iLen = 1
		} else if origin[i] == origin[i+1] {
			iLen++
		} else {
			result += strconv.Itoa(iLen) + origin[i:i+1]
			iLen = 1
		}
	}

	return result
}
