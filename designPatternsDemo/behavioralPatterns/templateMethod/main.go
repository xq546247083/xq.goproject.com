package main

import (
	"fmt"
)

func main() {
	basketball := &basketball{}
	play(basketball)

	football := &football{}
	play(football)
}

// ----------------游戏----------------
//  游戏
type IGame interface {
	init()
	startPlay()
	endPlay()
}

func play(game IGame) {
	game.init()
	game.startPlay()
	game.endPlay()
}

// ------------游戏的实现-----------------

// 篮球
type basketball struct {
}

func (this *basketball) init() {
	fmt.Println("篮球初始化")
}

func (this *basketball) startPlay() {
	fmt.Println("开始玩篮球")
}

func (this *basketball) endPlay() {
	fmt.Println("玩美滋滋了")
}

// 足球
type football struct {
}

func (this *football) init() {
	fmt.Println("足球初始化")
}

func (this *football) startPlay() {
	fmt.Println("开始玩足球")
}

func (this *football) endPlay() {
	fmt.Println("玩美滋滋了")
}
