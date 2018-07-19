// 1、普通方法调用只依赖于类型签名，也就是类名。
// 2、*T 的方法集包含了所有以 *T 或 T 为 receiver 的方法集。
// 导致了，如果*T实现了接口，不代表T也实现了该接口。但是T实现了接口，那么*T就一定实现了该接口。

package main

import (
	"fmt"
)

func main() {
	var p1 iPerson
	var p2 iPerson

	p1 = person{
		name: "xq",
		age:  11,
	}
	p2 = &person{}

	p1.name2()
	p1.age2()

	p2.name2()
	p2.age2()
}

type person struct {
	name string
	age  int
}

func (p person) name2() {
	fmt.Println(p.name)
}

func (p *person) age2() {
	fmt.Println(p.age)
}

type iPerson interface {
	name2()
	age2()
}
