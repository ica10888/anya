package routers

import (
	"anya/controllers"
	"github.com/astaxie/beego"
)

func init() {

	/*固定路由*/
    beego.Router("/", &controllers.MainController{})

	beego.BConfig.EnableGzip = true
	beego.BConfig.RouterCaseSensitive = true
	beego.BConfig.MaxMemory = 1 << 26

	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.CopyRequestBody = true

	/*设置模板取值方式*/
	//beego.BConfig.WebConfig.TemplateLeft = "${"
	//beego.BConfig.WebConfig.TemplateRight = "}"
	/*页面文件路径*/
	//beego.BConfig.WebConfig.ViewsPath="views/user"
	/*注解路由*/
	beego.Include(&controllers.DocListController{})
	beego.Include(&controllers.DocController{})
}
