package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dtNow := time.Now()

	rand.Seed(time.Now().Unix())
	for i := 1; i <= 100000; i++ {
		// rand.Seed(time.Now().Unix())
		rand.Intn(100)
	}

	fmt.Println(time.Now().Sub(dtNow).Seconds())
}
