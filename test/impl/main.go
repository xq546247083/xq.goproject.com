package main

import (
	"encoding/json"
	"fmt"
)

// 测试了继承相关的一些东西
func main() {
	a := new(Person)
	a.Base.Name = "1"
	a.Name = "2"

	// 匿名的组合表现
	fmt.Println(a.Base.Name, a.Name)
	a.Base.hello()
	a.hello()

	// 这里子类的name没了
	aByte, _ := json.Marshal(a)
	fmt.Println(string(aByte))
}

type Base struct {
	Name string
}

type Person struct {
	Base // 这种就是组合，只是是一个匿名的组合而已
	Name string
	Age  int32
}

func (this *Base) hello() {
	fmt.Println(this.Name)
}

func (this *Person) hello() {
	fmt.Println(this.Name)
}

// -------------接口才会继承-----------------
// 基础接口
type BaseImpl interface {
	hello()
}

// 测试继承，person继承了BaseImpl接口，所以返回person不报错
func testImpl() BaseImpl {
	return new(Person)
}
