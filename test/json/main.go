package main

import (
	"fmt"
    "encoding/json"
)

func main() {
    a:=new(Base)
    a.Name="xxx"
    str2,_:=json.Marshal(a)
    fmt.Println(string(str2))
    
    // tag的使用
    str:="{\"testTag\":\"temp\"}"
    json.Unmarshal([]byte(str),a)

    fmt.Println(a)
}

type Base struct{
    // tag的使用
    Name string `json:"testTag"`
}