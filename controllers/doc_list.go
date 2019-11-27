package controllers

import (
	"anya/models"
	"github.com/astaxie/beego"
)

type DocListController struct {
	beego.Controller
}

func (this *DocListController) URLMapping() {
	this.Mapping("GET", this.GetDocList)
}

// @router /doc/docList [get]
func (this *DocListController) GetDocList() {
	this.Data["json"] = &models.DescMaps
	this.ServeJSON()
}
