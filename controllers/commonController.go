// Package controllers
package controllers

import (
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

func (cC *CommonController) Get() {
	cC.Data["json"] = "Version: 1.2"
	cC.ServeFormatted()
}
