package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 6}
	b := 0
	fmt.Println(searchInsert(a, b))
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		nums = append(nums, target)
		return 0
	}

	if target < nums[0] {
		nums = append([]int{target}, nums...)
		return 0
	}

	for i := 0; i < length; i++ {
		if target == nums[i] {
			return i
		} else {
			if i == length-1 {
				nums = append(nums, target)
				return length
			} else if nums[i] < target && target < nums[i+1] {
				nums = append(nums[:i+1], append([]int{target}, nums[i+1:]...)...)
				return i + 1
			}
		}
	}

	return 0
}
