package main

import (
	"demo_login/controllersb"
	"github.com/astaxie/beego"
)


func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Router("/", &controllersb.IndexController{})
	beego.Router("/index", &controllersb.RegistController{})
	beego.Router("/query", &controllersb.UserController{})
	beego.Run()
}
