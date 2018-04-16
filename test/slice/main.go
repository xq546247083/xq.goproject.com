package main

import (
	"fmt"
	
)

func main(){
	a:=[]int{}
	a1:=make([]int,0)
	a2:=make([]int,0,0)

	a=append(a,1)
	a1=append(a1,1)
	a2=append(a2,1)
	fmt.Println(a,a1,a2)
}