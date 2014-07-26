package routers

import (
	"github.com/astaxie/beego"
	"short_url/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
