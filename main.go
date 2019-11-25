package main

import (
	"anya/models"
	_ "anya/routers"
	"github.com/astaxie/beego"
)

func main() {
	defer models.DataBase.Close()
	models.Init()
	go models.HandleSchedule()
	beego.Run()
}
