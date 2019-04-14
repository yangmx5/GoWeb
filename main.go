package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/yangmx5/goweb/routers"
)

func main() {
	fmt.Println("hello world")
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.Run()
}
