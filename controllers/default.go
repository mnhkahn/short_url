package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"short_url/models"
	"strings"
)

var (
	HTTP_HEAD = "http://"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) Shorten() {
	url := models.Url{}
	if err := this.ParseForm(&url); err != nil {
		//handle error
	}
	url.Shorten()

	url.ShortUrl = HTTP_HEAD + beego.AppConfig.String("default.domain") + url.ShortUrl
	if strings.Index(url.LongUrl, HTTP_HEAD) == -1 {
		url.LongUrl = HTTP_HEAD + url.LongUrl
	}
	this.Data["LongUrl"] = url.LongUrl
	this.Data["ShortUrl"] = url.ShortUrl
	this.TplNames = "index.tpl"
}

func (this *MainController) Lengthen() {
	url := models.Url{}
	if err := this.ParseForm(&url); err != nil {
		//handle error
	}
	if strings.Index(url.ShortUrl, HTTP_HEAD) != -1 {
		url.ShortUrl = url.ShortUrl[strings.LastIndex(url.ShortUrl, "/"):]
	}

	url.Lengthen()

	if strings.Index(url.LongUrl, HTTP_HEAD) == -1 {
		url.LongUrl = HTTP_HEAD + url.LongUrl
	}
	if strings.Index(url.ShortUrl, beego.AppConfig.String("default.domain")) == -1 {
		url.ShortUrl = HTTP_HEAD + beego.AppConfig.String("default.domain") + url.ShortUrl
	}
	this.Data["LongUrl"] = url.LongUrl
	this.Data["ShortUrl"] = url.ShortUrl
	this.TplNames = "index.tpl"
}

func (this *MainController) Redirect() {
	url := models.Url{}
	url.ShortUrl = this.Ctx.Input.Param(":short_url")
	url.Lengthen()
	if strings.Index(url.LongUrl, HTTP_HEAD) == -1 {
		url.LongUrl = HTTP_HEAD + url.LongUrl
	}
	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, url.LongUrl, http.StatusMovedPermanently)
}
