package main

import (
	"sync"
	
)

var locker sync.Mutex
var locker1 sync.Mutex

func main() {
    lock();
}

func lock(){
    locker.Lock()
    defer locker.Unlock()

    lock1()
}

func lock1(){
    locker1.Lock()
    defer locker1.Unlock()
}