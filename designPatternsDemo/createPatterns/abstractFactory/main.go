package main

import (
	"fmt"
)

func main() {
	factoryProducerTemp := new(factoryProducer)
	pizzaFactory := factoryProducerTemp.getFactory("Pizza")
	pizza := pizzaFactory.getPizza("ChickenPizza")
	pizza.Eat()

	colorFactory := factoryProducerTemp.getFactory("Color")
	color := colorFactory.getColor()
	color.Set()
}

//--------------------工厂创造器--------------------
type factoryProducer struct{}

// 获取工厂
func (this *factoryProducer) getFactory(name string) abstractFactory {
	switch name {
	case "Pizza":
		return new(pizzaFactory)
	default:
		return new(colorFactory)
	}
}

//--------------------抽象工厂--------------------
// 定义了一个抽象工厂
type abstractFactory interface {
	getPizza(name string) Pizza
	getColor() Color
}

//--------------------披萨工厂--------------------
// 披萨工厂
type pizzaFactory struct{}

// 去披萨工厂获取披萨
func (this *pizzaFactory) getPizza(name string) Pizza {
	switch name {
	case "ChickenPizza":
		return new(ChickenPizza)
	default:
		return new(CheesePizza)
	}
}

// 去披萨工厂上色
func (this *pizzaFactory) getColor() Color {
	return nil
}

//--------------------颜色工厂--------------------
// 颜色工厂
type colorFactory struct{}

// 去颜色工厂获取披萨
func (this *colorFactory) getPizza(name string) Pizza {
	return nil
}

// 去披萨工厂上色
func (this *colorFactory) getColor() Color {
	return new(Red)
}

//--------------------披萨--------------------
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

//--------------------颜色--------------------
// 颜色
type Color interface {
	Set()
}

type Red struct {
}

func (this *Red) Set() {
	fmt.Println("set Red")
}
