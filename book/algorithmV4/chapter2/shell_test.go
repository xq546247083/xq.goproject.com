// 以3*stepNum+1的步子划分数组，然后把下一个区域的数据和上一个区域的数据比较，把小的数据交换到上一个区域。
// stepNum逐步减少，当stepNum为1的时候，本质就是插入排序，因为之前部分交换过数据，这样可以减少交换的次数。
// 这就是优化过的插入排序。
package chapter1

import "testing"

// TestShellSort 测试
func TestShellSort(t *testing.T) {
	nums := getRandomList(10000)
	t.Errorf("%v", shellSort(nums))
}

// 希尔排序
func shellSort(nums []int) []int {
	lenNum := len(nums)

	// stepNum 1 4 13 40 12 362 1093...
	stepNum := 1
	for {
		if stepNum >= lenNum/3 {
			break
		}

		stepNum = 3*stepNum + 1
	}

	// 循环，分片给数组排序
	for {
		if stepNum < 1 {
			break
		}

		for i := stepNum; i < lenNum; i++ {
			for j := i; j >= stepNum && less(nums[j], nums[j-stepNum]); j -= stepNum {
				nums = exch(nums, j, j-stepNum)
			}
		}

		// 挑战步长
		stepNum = stepNum / 3
	}

	return nums
}
