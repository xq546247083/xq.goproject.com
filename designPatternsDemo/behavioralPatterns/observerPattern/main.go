package main

import (
	"fmt"
	"sync"
)

func main() {
	// 新建对象
	subjectObj := &subject{}

	observerAObj := &observerA{}
	observerBObj := &observerB{}

	// 注册方法
	subjectObj.registerFunc(observerAObj.hello)
	subjectObj.registerFunc(observerAObj.hello)
	subjectObj.registerFunc(observerBObj.hi)

	subjectObj.do("哈哈")
}

// ----------------被观察者----------------

//  被观察者类
type subject struct {
	listFunc []func(string)
	lock     sync.RWMutex
}

func (this *subject) registerFunc(funcObj func(string)) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.listFunc = append(this.listFunc, funcObj)
}

// 调用方法
func (this *subject) do(str string) {
	this.lock.Lock()
	defer this.lock.Unlock()

	for _, item := range this.listFunc {
		item(str)
	}
}

// ------------观察者-----------------

//  观察者A
type observerA struct {
}

func (this *observerA) hello(str string) {
	fmt.Println("A,hello:", str)
}

//  观察者B
type observerB struct {
}

func (this *observerB) hi(str string) {
	fmt.Println("A,hi:", str)
}
