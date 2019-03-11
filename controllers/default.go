package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "yangmx5.me"
	c.Data["Email"] = "yangmx@gmail.com"
	c.TplName = "index.tpl"
}
