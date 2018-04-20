package main

import (
	"fmt"
)

func main(){
	// 基本的咖啡
	smipleCoffeeTemp:=new(smipleCoffee)
	
	// 在基本的咖啡上加牛奶
	milkCoffeeTemp:=new(milkCoffee)
	milkCoffeeTemp.BaseCoffee=smipleCoffeeTemp
	fmt.Println(milkCoffeeTemp.getDesc(),"----",milkCoffeeTemp.getPrice())

	// 在基本的咖啡上加糖
	sugarCoffeeTemp:=new(sugarCoffee)
	sugarCoffeeTemp.BaseCoffee=smipleCoffeeTemp
	fmt.Println(sugarCoffeeTemp.getDesc(),"----",sugarCoffeeTemp.getPrice())

	// 在牛奶咖啡上加糖（装饰者模式的重点在于此处，在任何咖啡上，可以任意组装咖啡）
	sugarMilkCoffeeTemp:=new(sugarCoffee)
	sugarMilkCoffeeTemp.BaseCoffee=milkCoffeeTemp
	fmt.Println(sugarMilkCoffeeTemp.getDesc(),"----",sugarMilkCoffeeTemp.getPrice())

	// 在牛奶咖啡上加双倍糖
	sugarSugarMilkCoffeeTemp:=new(sugarCoffee)
	sugarSugarMilkCoffeeTemp.BaseCoffee=sugarMilkCoffeeTemp
	fmt.Println(sugarSugarMilkCoffeeTemp.getDesc(),"----",sugarSugarMilkCoffeeTemp.getPrice())
}

// -----------咖啡接口，以及普通咖啡的实现-----------

// 咖啡接口
type iCoffee interface{
	getDesc()string
	getPrice()int32
}

// 普通咖啡
type smipleCoffee struct{}

func (this *smipleCoffee)getDesc()string{
	return	fmt.Sprintf("咖啡")
}

func (this *smipleCoffee)getPrice()int32{
	return 5
}

// -------------------下面开始装饰者,组合了咖啡对象--------------------------

// 装饰者
type decorator struct{
	BaseCoffee iCoffee 
}

func (this *decorator)getDesc()string{
	return this.BaseCoffee.getDesc()
}

func (this *decorator)getPrice()int32{
	return this.BaseCoffee.getPrice()
}


// -------------------装饰着的具体的实现，对其具体方法进行了扩展--------------------------

// 牛奶咖啡
type milkCoffee struct{
	decorator
}

func (this *milkCoffee)getDesc()string{
	return this.decorator.getDesc()+"加牛奶"
}

func (this *milkCoffee)getPrice()int32{
	return this.decorator.getPrice()+3
}


// 加糖咖啡
type sugarCoffee struct{
	decorator
}

func (this *sugarCoffee)getDesc()string{
	return this.decorator.getDesc()+"加糖"
}

func (this *sugarCoffee)getPrice()int32{
	return this.decorator.getPrice()+1
}
