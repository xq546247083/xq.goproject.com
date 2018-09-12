package main

import "fmt"

func main() {
	fmt.Println(mySqrt(101))
}

func mySqrt(x int) int {
	i := 1
	for {
		if i*i == x {
			break
		}

		if i*i > x {
			i = i - 1
			break
		}

		i++
	}

	return i
}
