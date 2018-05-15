package main

import (
	"fmt"
)

// 测试不允许copy的提示的结构体构建
type noCopy struct{}

func (*noCopy) Lock() {}

type student struct {
	noCopy noCopy
	name   string
	age    int32
}

func main() {
	var student1 student
	student1.name = "x"
	PrintName(student1)
}

// 一般来说，这里出现不允许copy的提示
func PrintName(studentTemp student) {
	fmt.Println(studentTemp.name)
}
