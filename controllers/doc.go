package controllers

import (
	"encoding/json"
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
func (d *DocController) GetDoc() {
	strTittle := d.Ctx.Input.Param(":tittle")
	logs.Info(">>>> Tittle: %s <<<<",strTittle)
	d.Data["Tittle"] = strTittle
	bytes,err := json.Marshal(strTittle)
	if err != nil {
		panic("json Marshal failed")
	}
	d.Ctx.ResponseWriter.Write(bytes)
}