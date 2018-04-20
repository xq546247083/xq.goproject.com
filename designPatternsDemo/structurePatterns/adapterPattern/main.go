package main

import (
	"fmt"
)

func main(){
	concreteTargetTemp:=new(concreteTarget)
	concreteTargetTemp.request()

	adapterTemp:=new(adapter)
	adapterTemp.request()
}

//  标准接口
type target interface{
	request()
}

// 标准实现
type concreteTarget struct{}

func (this *concreteTarget)request(){
	fmt.Println("一个标准的实现");
}

// --------------------注意，适配器就是为了下面的这种类服务的-----------------------------

// 一个特殊的结构
type specialTarger struct{}

func (this *specialTarger)hello(){
	fmt.Println("这是一个特殊的请求，Hello");
}

// -----------注意，这就是适配器类了-----------

type adapter struct{
	specialTarger
}

// 用通用的请求，调用特殊的请求，以达到适配的效果
func (this *adapter)request(){
	this.specialTarger.hello();
}