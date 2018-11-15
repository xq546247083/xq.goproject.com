package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("ababccb"))
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	sLen := len(s)
	tempMap := make(map[uint8]int, sLen)

	// i为标记子串的起始位置
	for j, i := 0, 0; j < sLen; j++ {
		// 当发现重复字符时，重置子串的起始位置
		if index, ok := tempMap[s[j]]; ok {
			i = max(index, i)
		}

		fmt.Println("a", i)
		// 计算当前子序列的长度
		result = max(result, j-i+1)
		tempMap[s[j]] = j + 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
