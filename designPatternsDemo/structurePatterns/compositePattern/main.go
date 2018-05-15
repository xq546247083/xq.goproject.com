package main

import (
	"fmt"
)

func main() {
	// 根节点
	root := &composite{
		name: "root",
	}

	// 2个节点
	composite1 := &composite{
		name: "composite1",
	}

	composite2 := &composite{
		name: "composite2",
	}

	// 添加2个节点
	root.add(composite1)
	root.add(composite2)

	// 叶子
	leaf := &leaf{
		name: "leaf",
	}

	// 节点1添加叶子
	composite1.add(leaf)

	for _, child := range root.childList {
		child.printName()
		compositeTemp, status := child.(*composite)
		if status {
			for _, compositeTemp2 := range compositeTemp.childList {
				compositeTemp2.printName()
			}
		}
	}
}

// -----------组件接口-----------

// 组件接口
type iComponent interface {
	add(iComponent)
	remove(iComponent)
	printName()
}

// -----------组合节点-----------

// 组合节点
type composite struct {
	name      string
	childList []iComponent
}

func (this *composite) add(component iComponent) {
	this.childList = append(this.childList, component)
}

func (this *composite) remove(component iComponent) {
	for index, item := range this.childList {
		if item == component {
			this.childList = append(this.childList[:index], this.childList[index+1:]...)
			break
		}
	}
}

func (this *composite) printName() {
	fmt.Println(this.name)
}

// -------------------叶子--------------------------

// 组合节点
type leaf struct {
	name string
}

// 叶子没有增加或者减少节点
func (this *leaf) add(component iComponent) {
}

func (this *leaf) remove(component iComponent) {
}

func (this *leaf) printName() {
	fmt.Println(this.name)
}
