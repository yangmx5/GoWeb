package routers

import (
	"github.com/astaxie/beego"
	"github.com/yangmx5/goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/users",
		beego.NSRouter("/", &controllers.UserController{}),
		beego.NSRouter("/register", &controllers.UserController{}, "get:Register"),
		beego.NSRouter("/login", &controllers.UserController{}, "get:Login"),
		beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
		beego.NSRouter("/passwd", &controllers.UserController{}, "post:Passwd"),
		beego.NSRouter("/uploads", &controllers.UserController{}, "post:Uploads"),
		beego.NSRouter("/downloads", &controllers.UserController{}, "get:Downloads"),
	)
	beego.AddNamespace(ns)
}
