package main

import (
	"fmt"
    "encoding/json"
)

func main() {
    a:=new(Person)
    a.Base.Name="1"
    a.Name="2"

    // go的继承不是真正的继承，而是一个匿名的组合
    fmt.Println(a.Base.Name,a.Name)
    a.Base.hello()
    a.hello()

    // 这里子类的name没了
    aByte,_:=json.Marshal(a)
    fmt.Println(string(aByte))
}

type Base struct{
    Name string 
}

type Person struct{
    Base
    Name string 
    Age int32
}

func (this *Base)hello(){
    fmt.Println(this.Name)
}

func (this *Person)hello(){
    fmt.Println(this.Name)
}