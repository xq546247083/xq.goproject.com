package Chapter1

//二分法查找
//切片s是升序的
//k为待查找的整数
//如果查到有就返回对应角标,
//没有就返回-1
func BinarySearch(nums []int, k int) int {
	index, lenSlice := 0, len(nums)-1
	for index <= lenSlice {
		m := (index + lenSlice) >> 1
		if nums[m] < k {
			index = m + 1
		} else if nums[m] > k {
			lenSlice = m - 1
		} else {
			return m
		}
	}

	return -1
}
