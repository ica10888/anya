package controllers

import (
	"anya/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type DocController struct {
	beego.Controller
}

func (this *DocController) URLMapping() {
	this.Mapping("GET",this.GetDoc)
}

// @router /doc/:title [get]
func (this *DocController) GetDoc() {
	strTitle := this.Ctx.Input.Param(":title")
	logs.Info("%s",strTitle)
	this.Data["Title"] = strTitle
	this.Data["Content"] = models.Markdowns[strTitle].DocContent
	this.TplName = "page.tpl"
}