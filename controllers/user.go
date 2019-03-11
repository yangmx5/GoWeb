package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["json"] = map[string]string{
		"user": "get",
	}
	c.ServeJSON()
}

func (c *UserController) Register() {
	c.Data["json"] = map[string]string{
		"hello": "me",
	}
	c.ServeJSON()
}

func (c *UserController) Login() {
	c.ServeJSON()
}

func (c *UserController) Logout() {
	c.ServeJSON()
}

func (c *UserController) Passwd() {
	c.ServeJSON()
}

func (c *UserController) Uploads()   {}
func (c *UserController) Downloads() {}
