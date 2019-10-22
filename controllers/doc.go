package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type DocController struct {
	beego.Controller
}

func (d *DocController) URLMapping() {
	d.Mapping("GET",d.GetDoc)
}

// @router /doc/:tittle [get]
func (c *DocController) GetDoc() {
	strTittle := c.Ctx.Input.Param(":tittle")
	logs.Info(">>>> Tittle: %s <<<<",strTittle)
	c.Data["Tittle"] = strTittle
	c.TplName = "page.tpl"
}