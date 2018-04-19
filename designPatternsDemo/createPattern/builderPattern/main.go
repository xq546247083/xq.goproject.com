package main

import (
	"fmt"
)

func main(){
	// 创建导演
	directorObj:=new(director)

	// 创建相应的构建者
	builderObj:=new(builder)

	// 调用构建方法
	directorObj.construct(builderObj)

	// 获取构建出来的结果
	product:=builderObj.GetResult()

	fmt.Println(product)	
}

// -----------------------导演-----------------------

type director struct{}

// 构造
func (this *director) construct(builderObj *builder){
	builderObj.buildPartA()
	builderObj.buildPartB()
}

// -----------------------建设者-----------------------

// 定义一个建设着
type iBuilder interface{
	buildPartA()
	buildPartB()
}

// 建设者实现
type builder struct{
	productTemp product
}

// 建设部分A
func (this *builder) buildPartA(){
	this.productTemp.partA="partA"
}

// 建设部分B
func (this *builder) buildPartB(){
	this.productTemp.partB="partB"
}

// 获取产品
func (this *builder) GetResult() product{
	return this.productTemp
}

// -----------------------产品-----------------------

// 产品
type product struct{
	partA string
	partB string
}