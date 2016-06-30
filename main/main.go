package main

import "github.com/astaxie/beego"

type MainRouterController struct {
	beego.Controller
}

func main() {
	beego.Router("/", &MainRouterController{})
	beego.Run();
}

func (this *MainRouterController)Get() {
	this.Ctx.WriteString("hello, this my first go web app ~ ~")
}
