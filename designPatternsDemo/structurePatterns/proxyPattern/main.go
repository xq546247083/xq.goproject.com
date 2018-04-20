package main

import (
	"fmt"
)

func main(){
	proxyImageObj:=new(proxyImage)
	proxyImageObj.name="纹理图"
	proxyImageObj.display()
}

//  图片
type image interface{
	display()
}

// 真正的图片
type realImgae struct{
	name string 
}

func (this *realImgae)display(){
	fmt.Println("显示一个图片:",this.name);
}

// --------------------下面是代理类-----------------------------

// 代理图片
type proxyImage struct{
	realImgaeObj *realImgae
	name string 
}

func (this *proxyImage)display(){
	if this.realImgaeObj==nil{
		this.realImgaeObj=new(realImgae)
		this.realImgaeObj.name=this.name
	}
	
	this.realImgaeObj.display()
}