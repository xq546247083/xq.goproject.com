package main

import (
	"fmt"

)

func main() {
    testNewInterface();
}

func testNewInterface(){
    x:=&Person{Name:"xx"}
    x1:=new(Person)

     // 值类型复制就是一个完全的新对象（对应c#的深拷贝）
     x3:=*x
     x3.Name="x1"
     _=x3

    fmt.Println(x)
    fmt.Println(x1)

    // x2会nil，但是并不会报错，因为go的对象方法为一种关系，并不是指定类的方法（相当于第一个参数是调用对象）
    var x2 *Person
    x2.printHello()
}

func testRune(){
    var a rune
    fmt.Println(a)
}

func testVar(){
    // var一个值类型会自动初始化默认值
    var a Person

    // var一个指针，会初始化一个空指针
    var a1 *Person

    fmt.Println("xxx",a.Name);
    fmt.Println("xxx",a1.Name);
}

func (this *Person)printHello(){
    fmt.Println("hello")
}

type Person struct{
    Name string 
    Age int32
}

