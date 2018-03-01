package main

import (
	"fmt"
	
)

func main(){

	dasd:=make([]int32,10,20)
	a1:=dasd[2:3:10]
	_=a1

	a:=10
	f:=func() int{a=a*2;return 5}
	x:=[]int{a,f()}
	
	for a := 0; a < 100; a++ {
		fmt.Println(x)
	}	
}

type Encipher func(plaintext string) []byte

// encipher 作为一个方法类型
// 返回一个func
func GetItem(encipher Encipher) func(string)string{
	return func(dd string) string {
		return fmt.Sprintf("%x",encipher(dd));
	}
}