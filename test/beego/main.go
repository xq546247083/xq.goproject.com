package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	_ "xq.goproject.com/test/beego/routers"
)

func main() {
	beego.Router("/Test", &MainController{})
	beego.Get("Test1", func(c *context.Context) {
		c.Output.Body([]byte("bob"))
	})

	beego.Run()
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello")
}
