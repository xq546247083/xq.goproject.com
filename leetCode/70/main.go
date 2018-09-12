package main

import "fmt"

func main() {
	fmt.Println(climbStairs(6))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		// 前面2位的数字
		p2 := 1
		p1 := 2
		for i := 3; i < n; i++ {
			temp := p2
			p2 = p1
			p1 = p1 + temp
		}

		return p1 + p2
	}

}
