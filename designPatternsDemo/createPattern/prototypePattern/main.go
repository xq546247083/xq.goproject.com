package main

import (
	"fmt"
)

func main(){
	// 新建了一个人物A
	personA:=new(person)
	personA.name="A"
	personA.age=1

	// 克隆一个人物B
	personB:=personA.clone()
	// 修改人物A的名字为D
	personA.name="D"

	fmt.Println(personA)
	// 人物B的名字不变化
	fmt.Println(personB)
}

type iPrototype interface{
	clone() iPrototype
}

type person struct{
	name string 
	age int32
}

func (this *person)clone() person{
	return *this
}