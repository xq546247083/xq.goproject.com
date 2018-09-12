package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("x    dasdas"))
}

func lengthOfLastWord(s string) int {
	oldTemp := ""
	temp := ""
	for _, v := range s {
		if string(v) == " " {
			if temp != "" {
				oldTemp = temp
			}

			temp = ""
		} else {
			temp = temp + string(v)
		}
	}

	if temp != "" {
		return len(temp)
	}

	return len(oldTemp)
}
