// go的时间包里面，有2个时间，一个是系统时间，一个是递增时间。
// 系统时间和系统的时间是实时同步的。
// 递增时间却是按照时间流逝自己算的。

// 在时间计算时间差的时候（Sub），使用的是递增时间。

package main

import (
	"fmt"
	"time"
)

func main() {

	dtNow := time.Now()
	for {
		fmt.Printf("%s", time.Since(dtNow))
		fmt.Println("-------------------", time.Now().Unix()-dtNow.Unix())

		time.Sleep(time.Second)
	}
}
