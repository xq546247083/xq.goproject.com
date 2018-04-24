package main

import (
	"fmt"
)

func main(){
	// 车的对象
	busObj:=new(bus)
	carObj:=new(car)
	
	// 路的对象
	unpavedRoadObj:=new(unpavedRoad)
	cementRoadObj:=new(cementRoad)

	// 通过组合的方式，把路桥接过来，并可以动态的修改路
	// 跑起来
	busObj.setRoad(unpavedRoadObj)
	busObj.driveOnRoad()

	// 跑起来
	busObj.setRoad(cementRoadObj)
	busObj.driveOnRoad()

	// 跑起来
	carObj.setRoad(cementRoadObj)
	carObj.driveOnRoad()
}

// ----------------桥接模式的【桥】----------------

// 路的接口
type iRoad interface{
	drive()
}

// ----------------抽象路这边的对象----------------

// 石头路
type unpavedRoad struct{}

func (this *unpavedRoad)drive(){
	fmt.Println("行驶在石头路上");
}

// 水泥路
type cementRoad  struct{}

func (this *cementRoad)drive(){
	fmt.Println("行驶在水泥路上");
}

// --------------------抽象的车-----------------------------

// 车的接口
type iCehicle interface{
	driveOnRoad()
}

// 抽象的一辆车，组合了路的接口
type vehicle struct{
	roadObj iRoad
}

// 初始化车行驶的路的对象
func (this *vehicle)setRoad(roadObj iRoad){
	this.roadObj=roadObj
}

// 初始化车行驶的路的对象
func (this *vehicle)driveOnRoad(){
	panic("这是一条假的路，很危险啊!")
}


// ---------------------车的实现-----------------------

// 小汽车
type car struct{
	vehicle
}
 
// 初始化车行驶的路的对象
func (this *car)driveOnRoad(){
	fmt.Printf("小汽车");
	if this.roadObj!=nil{
		this.roadObj.drive();
	}else{
		panic("没路怎么跑，大兄弟!")
	}
}

// 公共汽车
type bus struct{
	vehicle
}
 
// 初始化车行驶的路的对象
func (this *bus)driveOnRoad(){
	fmt.Printf("公共汽车");
	if this.roadObj!=nil{
		this.roadObj.drive();
	}else{
		panic("没路怎么跑，大兄弟!")
	}
}