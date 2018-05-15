package main

import (
	"sync"
)

func main() {
	temp := GetInstance()
	_ = temp
}

// 定义一个类
type person struct{}

// 一个单例对象
var instance *person

// 运行一次的方法
var once sync.Once

func GetInstance() *person {
	once.Do(func() {
		instance = new(person)
	})

	return instance
}
