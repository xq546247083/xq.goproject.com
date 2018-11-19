// 归并两个有序的数组，为一个有序数组。
// 归并排序。
package chapter1

import (
	"testing"
)

// TestMergeSort2 测试
func TestMergeSort2(t *testing.T) {
	nums := getRandomList(10000000)
	t.Errorf("%v", mergeSort2(nums, 0, len(nums)-1))
}

// 优化的归并排序，在数据量大时，使用归并排序，其他时候，使用插入排序
func mergeSort2(nums []int, min, max int) []int {
	// 只有一个元素的时候，因为只有一个，那么它是肯定有序的，所以直接返回数组
	if min >= max {
		return nums
	}

	mid := min + (max-min)/2
	if max-min > 20 {
		// 递归排序左半边
		mergeSort(nums, min, mid)
		// 递归排序右半边
		mergeSort(nums, mid+1, max)
		// 合并左右的排序数据
		merge(nums, min, mid, max)
	} else {
		// 插入排序左半边
		insertSort2(nums, min, mid)
		// 插入排序右半边
		insertSort2(nums, mid+1, max)
		// 合并左右的排序数据
		merge(nums, min, mid, max)
	}

	return nums
}

// 插入排序
func insertSort2(nums []int, min, max int) []int {
	for i := min; i <= max; i++ {
		for j := i; j > min && less(nums[j], nums[j-1]); j-- {
			nums = exch(nums, j, j-1)
		}
	}

	return nums
}
