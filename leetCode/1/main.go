package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{-3, 4, 3, 90}, 0))
}

func twoSum(nums []int, target int) []int {
	staMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		result := target - nums[i]

		rIndex, status := staMap[result]
		if status {
			return []int{i, rIndex}
		} else {
			staMap[nums[i]] = i
		}
	}

	return []int{0, 0}
}
