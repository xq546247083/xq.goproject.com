package main

import (
	"fmt"
)

func main(){
	proxyImageObj:=new(proxyImage)
	proxyImageObj.name="纹理图"
	proxyImageObj.display()
}

// ----------------图片----------------
//  图片
type image interface{
	display()
	loadFromDisk()
}

// 真正的图片
type realImgae struct{
	name string 
}

func (this *realImgae)loadFromDisk(){
	fmt.Println("加载:",this.name);
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

// 代理模式组合了从硬盘加载，且第二次显示的时候，不需要再加载
// 而且代理图片少了一个加载的方法，控制了方法的数量
func (this *proxyImage)display(){
	if this.realImgaeObj==nil{
		this.realImgaeObj=new(realImgae)
		this.realImgaeObj.name=this.name
		this.realImgaeObj.loadFromDisk()
	}
	
	this.realImgaeObj.display()
}