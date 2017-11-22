// Package controllers
package controllers

import (
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

// @router /version [get]
func (cC *CommonController) GetVersion() {
	cC.Data["json"] = "version: 1.32"
	cC.ServeJSON()
}
