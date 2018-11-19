package chapter1

import "testing"

// 比较大小
func less(a, b int) bool {
	return a < b
}

// 交换数组元素
func exch(nums []int, a, b int) []int {
	nums[a], nums[b] = nums[b], nums[a]
	return nums
}

// 选择排序
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := min + 1; j < len(nums); j++ {
			if less(nums[j], nums[min]) {
				min = j
			}
		}

		nums = exch(nums, i, min)
	}

	return nums
}

// TestSelectSort 测试选择排序
func TestSelectSort(t *testing.T) {
	t.Errorf("%v", selectSort([]int{2, 5, 1, 4, 66, 3, 3, 2}))
}
