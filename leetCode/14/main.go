package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
}

func longestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) == 0 {
		return result
	}

	i := 0
	for {
		j := 0
		tempStr := ""
		for _, item := range strs {
			// 先判断长度时候够
			if len(item) > i {
				if j == 0 {
					tempStr = item[i : i+1]
				} else {
					if tempStr != item[i:i+1] {
						return result
					}
				}
			} else {
				// 如果不够，则说明可以返回了
				return result
			}

			j++
		}

		result = fmt.Sprintf("%s%s", result, tempStr)
		i++
	}
}
