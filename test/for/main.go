package main

import(
	"fmt"
	"time"
	"sync"
	"runtime"
)

func main(){
	printTestTwo()
}

// -------------------------Test1--------------------------------------

type field struct {
	name string
}

func print(p *field){
	fmt.Println(p.name)
}

func TestGo(){
	data := [3]field{{"one"},{"two"},{"three"}}

	// go在循环的过程中，对指针(&v)是复用,指向了v1,v2,v3，指针都是一样的。但是对象是拷贝出来的。
	for _,v := range data {
		fmt.Println(v)
		fmt.Println(fmt.Sprintf("%p",&v))
	}

	time.Sleep(3 * time.Second)
}

// -------------------------Test1--------------------------------------

func printTestTwo() {
	runtime.GOMAXPROCS(1)
	
    wg := sync.WaitGroup{}
	wg.Add(20)

	// A:均为输出10，B:从0~9输出(顺序不定)。 
	// 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。 故go func执行时，i的值始终是10。
	// 第二个go func中i是函数参数，与外部for中的i完全是两个变量。 尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("A: ", i)
            wg.Done()
        }()
	}
	
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("B: ", i)
            wg.Done()
        }(i)
	}
	
    wg.Wait()
}