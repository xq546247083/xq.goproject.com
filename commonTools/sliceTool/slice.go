package sliceTool

// InsertIntSlice 在数组中插入元素
func InsertIntSlice(numSilice []int, index, item int) []int {
	numSilice = append(numSilice[:index], append([]int{item}, numSilice[index:]...)...)
	return numSilice
}
