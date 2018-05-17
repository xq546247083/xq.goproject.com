package main

import (
	"fmt"
	"sync"
)

func main() {
	circlePoolObj := &circlePool{
		data: make(map[string]*circle),
	}

	circle1 := circlePoolObj.getCircle("green")
	circle1.draw()

	circle2 := circlePoolObj.getCircle("red")
	circle2.draw()

	circle1.x = 1
	circle1.y = 1
	circle1.draw()
}

// ----------------圆----------------

// 圆
type circle struct {
	color  string
	x      int
	y      int
	radius int
}

// 画
func (this *circle) draw() {
	fmt.Println(this)
}

// --------------------图片储存池-----------------------------

// 图片储存池
type circlePool struct {
	// 池里面的图片
	data map[string]*circle

	// 锁
	lock sync.RWMutex
}

// 获取圆
func (this *circlePool) getCircle(color string) *circle {
	this.lock.Lock()
	defer this.lock.Unlock()

	// 如果有，直接返回
	circleObj, status := this.data[color]
	if status {
		return circleObj
	}

	// 如果没有，那么新建一个，保存起来返回
	circleObj = &circle{
		color: color,
	}
	this.data[color] = circleObj

	return circleObj
}
