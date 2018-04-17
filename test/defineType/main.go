//********************************************
//自定义类型用法
//********************************************
package main

import (
	"fmt"
)

// 测试type的用法

// 学生类型
type Student struct{
    Name string
    Age int32
}

// 高中班级
type HighClass []*Student

// 高中班级2
type HighClass2 struct{
    StudentList []Student
}

// 定义了一个方法
type ToString func(interface{})string

func main() {
    d1:=make(HighClass,5)
    d1[0]=&Student{Age:1,Name:"xxx"}

    var studentToString ToString
    studentToString=GetName
    fmt.Println(studentToString(d1[0]))
}

func GetName(studentObj interface{}) string{
    return fmt.Sprintln(studentObj)
}