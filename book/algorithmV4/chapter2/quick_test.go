// 如果指定元素的左边和右边的都是有序的。
// 且指定元素的左边都小于指定元素的右边，那么数组是有序的。
// 快速排序。
package chapter1

import (
	"testing"
)

// TestQuickSort 测试
func TestQuickSort(t *testing.T) {
	nums := getRandomList(10000)
	t.Errorf("%v", quickSort(nums, 0, len(nums)-1))
}

// 快速排序
func quickSort(nums []int, min, max int) []int {
	// 只有一个元素的时候，因为只有一个，那么它是肯定有序的，所以直接返回数组
	if min >= max {
		return nums
	}

	// 切分数组，并返回切分位置
	nums, index := partition(nums, min, max)
	// 对切分两边的数据，进行快速排序
	quickSort(nums, min, index-1)
	quickSort(nums, index+1, max)

	return nums
}

// 合并数组的左右侧，这个数组的左右侧必须是有序的
// 返回一个有序的合并数组
func partition(nums []int, min, max int) ([]int, int) {
	i, j := min, max+1
	temp := nums[min]
	for {
		// 从头开始，往右循环，找到比temp这个临时元素大的元素
		for {
			i++
			if i == max {
				break
			}

			if less(temp, nums[i]) {
				break
			}
		}

		// 从尾开始，往左循环，找到比temp这个临时元素小的元素
		for {
			j--
			if j == min {
				break
			}

			if !less(temp, nums[j]) {
				break
			}
		}

		// 如果相遇了，那么终止循环
		if i >= j {
			break
		}

		nums = exch(nums, i, j)
	}

	// 最后，和临时元素，也就是min位置的元素，交换位置。
	nums = exch(nums, min, j)
	return nums, j
}
