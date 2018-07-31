package main

import (
	"fmt"
)

// 切片测试
func main() {
	testPtr()

	testF := make([]int, 7, 7)
	fmt.Println(testF, len(testF), cap(testF))

	testF = append(testF, 1)
	fmt.Println(testF, len(testF), cap(testF))

	testCap()
}

// 测试数组的指针
func testPtr() {
	a1 := make([]int, 0, 1)
	a1 = append(a1, 1)
	ptr := &a1[0]

	// 这里改值是改到了
	*ptr = 5
	fmt.Println(a1)

	// 这里数组扩容后，现在数组的地址发生了改变
	a1 = append(a1, 1)
	*ptr = 6
	fmt.Println(a1)
	fmt.Println(ptr, &a1[0])
}

func testCap() {
	a := make([]int, 3, 5)
	b := a[1:3:3]
	c := append(b, 1)

	// 扩展时，重新创建了一个底层数据，地址变化
	fmt.Println(fmt.Sprintf("%p----%p--------%v", b, c, c))

	d := make([]int, 3, 5)
	e := d[1:3:5]
	f := append(e, 1)

	fmt.Println(fmt.Sprintf("%p----%p--------%v", e, f, f))
}

func test() {
	a := [3]int{1}
	a1 := make([]int, 0)
	a2 := make([]int, 0, 0)

	a1 = append(a1, 1)
	a2 = append(a2, 1)
	fmt.Println(a, a1, a2)

	TestSet(a)
	fmt.Println(a, a1, a2)

	TestSet2(&a)
	fmt.Println(a, a1, a2)
}

// 传递数组
func TestSet(intArray [3]int) {
	intArray[0] = 2
}

// 传递数组的地址
func TestSet2(intArray *[3]int) {
	(*intArray)[0] = 3
}
