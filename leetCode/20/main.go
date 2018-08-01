package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("[]{}[{}]"))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	in := []string{"(", "{", "["}
	out := []string{")", "}", "]"}

	// 结果
	result := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {
		tempStr := s[i : i+1]

		inFlag := false
		for _, item := range in {
			if tempStr == item {
				inFlag = true
			}
		}

		// 如果是入的标签，那么追加元素
		if inFlag {
			result = append(result, tempStr)
		} else {
			// 下面开始出
			outFlag := false
			for _, item := range out {
				if tempStr == item {
					outFlag = true
				}
			}

			// 如果没在in的数组里，又没在out的数组，那么数据有错误，返回false
			if !outFlag {
				return false
			} else {
				// 如果这个时候没有元素可出了，那么返回false
				lenResult := len(result)
				if lenResult <= 0 {
					return false
				}
				// 如果匹配成功，那么删除最后一个字符
				if !isEquip(result[lenResult-1], tempStr) {
					return false
				} else {
					result = result[:len(result)-1]
				}
			}
		}
	}

	if len(result) > 0 {
		return false
	}

	return true
}

func isEquip(s1, s2 string) bool {
	if (s1 == "(" && s2 == ")") ||
		(s1 == "{" && s2 == "}") ||
		(s1 == "[" && s2 == "]") {
		return true
	}

	return false
}
