package controllers

import "github.com/astaxie/beego"

type LayuiCmsluyunController struct {
	beego.Controller
}

func (c *LayuiCmsluyunController) Get() {

	c.TplName = "layuiCmsluyun/index.html"
}
