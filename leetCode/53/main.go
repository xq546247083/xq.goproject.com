package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3}))
}

func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > result {
				result = sum
			}
		}
	}

	return result
}
