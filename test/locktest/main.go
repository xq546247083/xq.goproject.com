package main

import (
	"fmt"
	"sync"
	"time"

	"xq.goproject.com/test/locktest/lockUtilTest"
)

// 测试锁，以及其异常
func main() {
	locklocktest()
	locklocktest2()

	fmt.Scanln()
}

func locklocktest() {
	userLock := lockUtilTest.GetLock(1)
	userLock.Lock()
}

func locklocktest2() {
	userLock := lockUtilTest.GetLock(1)
	userLock.Lock()
}

// 测试锁
func locktest() {
	var locker sync.Mutex
	locker.Lock()

	for i := 1; i <= 3; i++ {
		go func(j int) {
			fmt.Println("lock:", j)
			locker.Lock()
			fmt.Println("unlock:", j)
		}(i)
	}

	time.Sleep(1 * time.Second)
	locker.Unlock()
}

// 给已解锁的锁解锁，触发异常
func unlocktest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var locker sync.Mutex
	locker.Lock()
	locker.Unlock()
	locker.Unlock()
}
