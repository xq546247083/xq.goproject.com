package sliceTool

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	nums2 := make([]int, 0, 5)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 5)
	nums2 = append(nums2, 6)

	nums2 = InsertIntSlice(nums2, 1, 0)
	fmt.Println(nums2)
	t.Errorf(fmt.Sprintf("1"))
}
