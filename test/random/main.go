package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(Generate_Randnum())
	}
}

func Generate_Randnum() int {
	//  rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)

	fmt.Printf("rand is %v\n", rnd)

	rnds := rand.Float64()
	fmt.Printf("rnds is %v\n", rnds)

	return rnd
}
