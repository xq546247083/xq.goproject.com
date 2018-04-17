package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	nameList:=[]string{
		"xq",
		"xiaoqiang",
	}

	this.Data["NameList"] = nameList
	this.Data["Website"] = "xiaohe.info"
	this.Data["Email"] = "test@gmail.com"
	this.TplName = "index.tpl"
}
