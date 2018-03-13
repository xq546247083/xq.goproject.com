package main

import (
	"fmt"
	
)

type noCopy struct{}
func (*noCopy) Lock() {}

type student struct{
    noCopy noCopy
    name string
    age int32
}

// 出现不允许copy的提示
func main() {
    var student1 student
    student1.name="x"
    PrintName(student1)
}

func PrintName(studentTemp student){
    fmt.Println(studentTemp.name)
}

