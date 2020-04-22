package controllersb

import (
	"demo_login/api"
	"demo_login/modelb"
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	sess := c.StartSession()
	username := sess.Get("name")
	//fmt.Println("get",username)
	if username == nil || username == "" {
		c.TplName = "login.tpl"
	} else {
		c.TplName = "success.tpl"
	}

}

func (c *IndexController) Post() {
	user := &modelb.UserLogin{}
	user.Name = c.GetString("name")
	//fmt.Println("post",user.Name)
	user.Pwd = c.GetString("pwd")
	//fmt.Println("post",user.Pwd)
	err := api.ValidateUser(user)
	if err == nil {
		c.SetSession("当前登录用户", "admin")
		fmt.Println("当前Session:", c.CruSession)
		c.TplName = "lists.html"
	} else {
		fmt.Println(err)
		c.TplName = "error.tpl"
	}
}
