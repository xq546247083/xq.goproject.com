// 测试内联优化
// 命令：go build --gcflags=-m
// 结果：
// .\main.go:15:19: inlining call to testInline.(*Person).GetName
// .\main.go:17:31: inlining call to reflect.TypeOf
// .\main.go:17:31: inlining call to reflect.toType
// .\main.go:17:31: reflect.Type(reflect.t·2) escapes to heap
// .\main.go:17:31: personObj escapes to heap
// .\main.go:14:18: new(testInline.Person) escapes to heap
// .\main.go:18:35: reflectType.NumMethod() escapes to heap
// .\main.go:17:31: main &reflect.i·2 does not escape
// .\main.go:18:13: main ... argument does not escape

// 虽然被内联优化了，但是方法集依然存在。

package main

import (
	"fmt"
	"reflect"

	"xq.goproject.com/test/inlineFunc/testInline"
)

func main() {
	personObj := new(testInline.Person)
	personObj.GetName("a")

	reflectType := reflect.TypeOf(personObj)
	fmt.Println(reflectType.NumMethod())
}
