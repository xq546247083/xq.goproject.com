// 堆排序。
package chapter1

import (
	"testing"
)

// TestHeadkSort 测试
func TestHeadkSort(t *testing.T) {
	nums := getRandomList(10000)
	t.Errorf("%v", heapSort(nums))
}

func heapSort(nums []int) []int {
	// 构造初始堆
	// 从第一个非叶子结点从下至上，从右至左调整结构
	for i := len(nums) / 2; i >= 0; i-- {
		heapAdjust(nums[i:])
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapAdjust(nums[:i])
	}

	return nums
}

// 调整大顶堆
func heapAdjust(nums []int) {
	e := 0
	temp := nums[0]

	for i := 1; i < len(nums); i = i*2 + 1 {
		if i < len(nums)-1 && nums[i] < nums[i+1] {
			i++
		}

		if nums[i] > temp {
			nums[e] = nums[i]
			e = i
		} else {
			break
		}
	}
	nums[e] = temp
}
