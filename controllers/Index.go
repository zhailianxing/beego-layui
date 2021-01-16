package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {

	user := c.GetSession("userInfo")
	if user != nil { //兼容
		c.Data["user"] = user
	}
	c.TplName = "index.html"
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//  测试我的session
	//testUser := c.GetSession("userInfo")
	//fmt.Println("MainController  user:", testUser)

	c.TplName = "main.html"
}
