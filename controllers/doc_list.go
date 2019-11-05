package controllers

import (
	"anya/models"
	"github.com/astaxie/beego"
)

type DocListController struct {
	beego.Controller
}

func (this *DocListController) URLMapping() {
	this.Mapping("GET",this.GetDocList)
}

// @router /doc/docList [get]
func (this *DocListController) GetDocList() {
	descMaps := make(map[string]models.DocDescription)
	for k, v := range models.Markdowns {
		descMaps[k] = v.DocDesc
	}
	this.Data["json"] = &descMaps
	this.ServeJSON()
}
