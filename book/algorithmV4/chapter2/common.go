package chapter1

// 比较大小
func less(a, b int) bool {
	return a < b
}

// 交换数组元素
func exch(nums []int, a, b int) []int {
	nums[a], nums[b] = nums[b], nums[a]
	return nums
}
