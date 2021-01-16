package routers

import (
	"baseLayui/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/main", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/layuiCmsluyun", &controllers.LayuiCmsluyunController{})

	beego.Router("/menuList", &controllers.MenuController{})

	beego.AutoRouter(&controllers.MenuController{})

}
