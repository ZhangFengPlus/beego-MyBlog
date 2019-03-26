package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_"beego-blog/routers"
	"beego-blog/models"
)


func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}


func main() {
	beego.Run()
}

