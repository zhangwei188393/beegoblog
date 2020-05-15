package main

import (
	_ "beegoblog/routers"
	"beegoblog/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

