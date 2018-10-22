package main

import "fmt"

func main() {
	nums1 := make([]int, 6, 6)
	nums1[0] = 1
	nums1[1] = 2
	nums1[2] = 3

	nums2 := make([]int, 0, 3)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 5)
	nums2 = append(nums2, 6)

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 设置初始值
	count := m + n - 1
	m--
	n--

	// 从大到小，给数组赋值大的值
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[count] = nums1[m]
			m--
		} else {
			nums1[count] = nums2[n]
			n--
		}
		count--
	}

	for n >= 0 {
		nums1[count] = nums2[n]
		count--
		n--
	}
}
