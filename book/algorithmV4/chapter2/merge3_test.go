// 归并两个有序的数组，为一个有序数组。
// 归并排序。
package chapter1

import (
	"testing"
)

// TestMergelSort3 测试
func TestMergelSort3(t *testing.T) {
	nums := getRandomList(10000)
	t.Errorf("%v", mergeSort3(nums))
}

// 归并排序自底向上的实现
func mergeSort3(nums []int) []int {
	lenNums := len(nums)

	// 步长 1 2 4 8...向上合并数组
	for step := 1; step < lenNums; step = step + step {
		for j := 0; j < lenNums-step; j = j + step + step {
			merge(nums, j, j+step-1, min(lenNums-1, j+step+step-1))
		}
	}

	return nums
}
