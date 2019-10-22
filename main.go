package main

import (
	"anya/models"
	_ "anya/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}

