package routers

import (
	"github.com/astaxie/beego"
	"short_url/controllers"
)

func init() {
	beego.Router("/short_url/", &controllers.MainController{}, "*:Get")
	beego.Router("/short_url/shorten", &controllers.MainController{}, "post:Shorten")
	beego.Router("/short_url/lengthen", &controllers.MainController{}, "post:Lengthen")
	beego.Router("/?:short_url", &controllers.MainController{}, "get:Redirect")
}
