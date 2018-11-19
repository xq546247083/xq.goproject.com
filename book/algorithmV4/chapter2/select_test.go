// 选择当前序列最小的元素，并把它和当前序列的第一个元素交换。
// 剩下的元素，成为下一个当前序列。
// 这样实现排序。
package chapter1

import "testing"

// TestSelectSort 测试选择排序
func TestSelectSort(t *testing.T) {
	t.Errorf("%v", selectSort([]int{2, 7, 5, 1, 4, 66, 3, 3, 2}))
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
