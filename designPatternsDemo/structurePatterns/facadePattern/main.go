package main

import (
	"fmt"
)

func main() {
	shapeMakerObj := new(shapeMaker)
	shapeMakerObj.draw()
}

// ----------------各种图形------------------------

type shape interface {
	draw()
}

// 矩形
type rectangle struct{}

func (this *rectangle) draw() {
	fmt.Println("draw rectangle")
}

// 正方形
type square struct{}

func (this *square) draw() {
	fmt.Println("draw square")
}

// 圆形
type circle struct{}

func (this *circle) draw() {
	fmt.Println("draw circle")
}

// --------------------外观模式--------------------

type shapeMaker struct {
	rectangleObj *rectangle
	squareObj    *square
	circleObj    *circle
}

// 初始化对象(隐藏了多个对象的初始化)
func (this *shapeMaker) init() {
	this.rectangleObj = new(rectangle)
	this.squareObj = new(square)
	this.circleObj = new(circle)
}

// 调用画的方法
func (this *shapeMaker) draw() {
	this.rectangleObj.draw()
	this.squareObj.draw()
	this.circleObj.draw()
}
