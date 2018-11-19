// 归并两个有序的数组，为一个有序数组。
// 归并排序。
package chapter1

import (
	"testing"
)

// TestMergeSort 测试
func TestMergeSort(t *testing.T) {
	nums := getRandomList(10000000)
	t.Errorf("%v", mergeSort(nums, 0, len(nums)-1))
}

// 归并排序
func mergeSort(nums []int, min, max int) []int {
	// 只有一个元素的时候，因为只有一个，那么它是肯定有序的，所以直接返回数组
	if min >= max {
		return nums
	}

	mid := min + (max-min)/2
	// 递归排序左半边
	mergeSort(nums, min, mid)
	// 递归排序右半边
	mergeSort(nums, mid+1, max)
	// 合并左右的排序数据
	merge(nums, min, mid, max)

	return nums
}

// 合并数组的左右侧，这个数组的左右侧必须是有序的
// 返回一个有序的合并数组
func merge(nums []int, min, mid, max int) []int {
	// 复制数据到临时数组
	tempNums := make([]int, max-min+1, max-min+1)
	for k := min; k <= max; k++ {
		tempNums[k-min] = nums[k]
	}

	// 从mid位置划分，取小的一方的数据，归并回原数组
	i, j := min, mid+1
	for k := min; k <= max; k++ {
		// 如果前半数组被完全归并，那么，当前元素直接复制后半的数据
		if i > mid {
			nums[k] = tempNums[j-min]
			j++
		} else if j > max {
			nums[k] = tempNums[i-min]
			i++
		} else if less(tempNums[j-min], tempNums[i-min]) {
			nums[k] = tempNums[j-min]
			j++
		} else {
			nums[k] = tempNums[i-min]
			i++
		}
	}

	return nums
}
