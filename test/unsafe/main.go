package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	A string
	B int
}

type B struct {
	X string
	Y int
}

func main() {
	a := new(A)
	a.A = "hello"
	a.B = 1

	b := (*B)(unsafe.Pointer(a))
	fmt.Println(fmt.Sprintf("%s", b))
}
