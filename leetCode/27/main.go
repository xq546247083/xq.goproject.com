package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 2, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length <= 0 {
		return length
	}

	index := 0
	for j := 0; j < length; j++ {
		if nums[j] != val {
			nums[index] = nums[j]
			index++
		}
	}

	return index
}
