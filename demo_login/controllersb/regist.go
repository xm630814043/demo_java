package controllersb

import (
	"demo_login/api"
	"demo_login/modelb"
	"fmt"
	"github.com/astaxie/beego"
)

type RegistController struct {
	beego.Controller
}

func (c *RegistController) Get() {
	c.TplName = "regist.tpl"
}

func (c *RegistController) Post() {
	var user = modelb.UserLogin{}
	//fmt.Println(inputs)
	user.Name = c.GetString("name")
	user.Pwd = c.GetString("pwd")
	err := api.SaveUser(&user)
	if err == nil {
		c.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		c.TplName = "error.tpl"
	}
}
