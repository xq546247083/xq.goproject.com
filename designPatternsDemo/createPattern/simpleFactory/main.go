package main

import (
	"fmt"
)

func main(){
	temp:=getPizza("ChickenPizza1")
	temp.Eat()
}

// 获取披萨，这个方法就是简单工厂模式的关键点
func getPizza(name string)Pizza{
	switch name{
	case "ChickenPizza":
		return new(ChickenPizza)
	default:
		return new(CheesePizza)
	}
}

// paizza接口
type Pizza interface{
	Eat()
}

// 鸡腿的
type ChickenPizza struct{
	Size int32 // 有自己的大小
}

func (this *ChickenPizza)Eat(){
	fmt.Println("eat ChickenPizza")
}

// 奶酪的
type CheesePizza struct{
	Size int32 // 有自己的大小
}

func (this *CheesePizza) Eat(){
	fmt.Println("eat CheesePizza")
}