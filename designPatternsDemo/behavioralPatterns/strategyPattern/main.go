package main

import (
	"fmt"
)

func main() {
	contextObj := &context{}

	contextObj.setOperate(&add{})
	fmt.Println(contextObj.exec(3, 2))

	contextObj.setOperate(&reduce{})
	fmt.Println(contextObj.exec(3, 2))

	contextObj.setOperate(&multiply{})
	fmt.Println(contextObj.exec(3, 2))
}

// ----------------策略接口----------------

//  策略接口
type iStrategy interface {
	do(int, int) int
}

// ------------实现-----------------

// 增加
type add struct {
}

func (this *add) do(a int, b int) int {
	return a + b
}

// 减少
type reduce struct {
}

func (this *reduce) do(a int, b int) int {
	return a - b
}

// 剩
type multiply struct {
}

func (this *multiply) do(a int, b int) int {
	return a * b
}

// -------------------策略类-------------------

// 具体执行者
type context struct {
	operate iStrategy
}

// 设置执行类
func (this *context) setOperate(operate iStrategy) {
	this.operate = operate
}

// 执行
func (this *context) exec(a int, b int) int {
	return this.operate.do(a, b)
}
