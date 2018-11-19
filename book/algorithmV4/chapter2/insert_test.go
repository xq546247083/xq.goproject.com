// 从第一个元素开始，把当前元素插入到当前元素左侧合适的位置，使当前元素的左侧一直有序。
// 以实现排序。
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

// TestInsertSort 测试
func TestInsertSort(t *testing.T) {
	t.Errorf("%v", insertSort([]int{2, 7, 5, 1, 4, 66, 3, 3, 2}))
}
