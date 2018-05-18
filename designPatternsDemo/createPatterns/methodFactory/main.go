package main

import (
	"fmt"
)

func main() {
	// 新建一个工厂
	var factoryTemp factory
	factoryTemp = new(localFactory)

	temp := factoryTemp.getPizza()
	temp.Eat()
}

// 定义一个工厂接口
type factory interface {
	getPizza() Pizza
}

// 本地工厂
type localFactory struct{}

// 外地工厂
type nonLocalFactory struct{}

// 去本地工厂获取披萨
func (this *localFactory) getPizza() Pizza {
	return new(ChickenPizza)
}

// 去本地工厂获取披萨
func (this *nonLocalFactory) getPizza() Pizza {
	return new(CheesePizza)
}

// paizza接口
type Pizza interface {
	Eat()
}

// 鸡腿的
type ChickenPizza struct {
}

func (this *ChickenPizza) Eat() {
	fmt.Println("eat ChickenPizza")
}

// 奶酪的
type CheesePizza struct {
}

func (this *CheesePizza) Eat() {
	fmt.Println("eat CheesePizza")
}
