package main

import "fmt"

func main() {

	fmt.Println(strStr("", "a"))
}

func strStr(haystack string, needle string) int {
	nLen := len(needle)
	if nLen == 0 {
		return 0
	}

	hLen := len(haystack)
	for i := 0; i < hLen; i++ {
		isEquip := true
		for j := 0; j < nLen; j++ {
			// 如果超过了查找字符串长度，直接返回
			if i+j >= hLen {
				return -1
			}

			// 如果不等，查找下一个
			if haystack[i+j] != needle[j] {
				isEquip = false
				break
			}
		}

		if isEquip {
			return i
		}
	}

	return -1
}
