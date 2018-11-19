package chapter1

import "testing"

// 插入排序
func insertSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0 && less(nums[j], nums[j-1]); j-- {
			nums = exch(nums, j, j-1)
		}
	}

	return nums
}

// TestInsertSortt 测试
func TestInsertSortt(t *testing.T) {
	t.Errorf("%v", selectSort([]int{2, 7, 5, 1, 4, 66, 3, 3, 2}))
}
