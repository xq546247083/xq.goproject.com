package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// 新建一个原子值
	aValue := new(atomic.Value)

	// 存储数据,再存在会被替换掉
	aValue.Store([]int32{1, 22, 2, 35})
	aValue.Store([]int32{1, 22, 2, 36})
	fmt.Println(aValue.Load())

	// 存储新的原子值
	anotherStore(aValue)
	fmt.Println(aValue.Load())
}

// 引用类型，会改变原有的原子值
func anotherStore(newValue *atomic.Value) {
	newValue.Store([]int32{2, 3213123, 2, 35})
}
