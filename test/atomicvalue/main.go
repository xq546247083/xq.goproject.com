package main

import (
	"fmt"
	"sync/atomic"
	
)


func main() {
    aValue:= new(atomic.Value)

    aValue.Store([]int32{1,2,2,3})
    aValue.Store([]int32{1,22,2,35})

    fmt.Println(aValue.Load())
    anotherStore(aValue)
    fmt.Println(aValue.Load())
}

func anotherStore(newValue *atomic.Value){
    newValue.Store([]int32{2,3213123,2,35})
}