package chapter1

import (
	"math/rand"
	"time"
)

// 比较大小
func less(a, b int) bool {
	return a < b
}

// 交换数组元素
func exch(nums []int, a, b int) []int {
	nums[a], nums[b] = nums[b], nums[a]
	return nums
}

func getRandomList(len int) []int {
	nums := make([]int, 0, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := 0
	for {
		randomInt := r.Intn(1000)
		nums = append(nums, randomInt)

		i++
		if i >= len {
			break
		}
	}

	return nums
}
