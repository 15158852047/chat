package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	i := this.GetSession("id")
	if i != nil {

		this.TplName = "index.html"
		return
	}
	this.TplName = "login.html"
}
