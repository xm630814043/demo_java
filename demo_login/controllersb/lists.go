package controllersb

import (
	"demo_login/api"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.TplName = "lists.html"
}

func (c *UserController) Post() {
	pageno,err:=c.GetInt("pageno")
	if err!=nil{
		fmt.Println(err)
	}
	userList :=api.QuerydateUser(1,pageno)
	//fmt.Println("查询全表", len(userList),userList)
	c.Data["json"]=map[string]interface{}{"PageSize":1,"Page":pageno,"DataList":userList}
	c.ServeJSON()
}




