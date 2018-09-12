package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if i == 0 && digits[i] == 9 {
			digits[i] = 0
			digits = append([]int{1}, digits...)
			break
		}

		if digits[i] < 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}

	return digits
}
