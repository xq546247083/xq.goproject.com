package main

import (
	"fmt"
	"sync"
	"time"

	"xq.goproject.com/test/goChanPool/src"
)

var (
	wg sync.WaitGroup
)

func main() {
	dispather := src.NewDispather(10)
	dispather.Run()

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(_wg *sync.WaitGroup) {
			defer func() {
				_wg.Done()
			}()

			dispather.AddJob(test)
		}(&wg)
	}

	wg.Wait()
}

func test(dtNow time.Time) {
	fmt.Println("xxxx", dtNow)
}
