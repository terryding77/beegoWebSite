// Package controllers provides ...
package controllers

import (
	"encoding/json"
	"goWebTest/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @router / [get]
func (uC *UserController) GetAllUsers() {
	uC.Data["json"] = models.GetAllUsers()
	uC.ServeJSON()
}

// @router / [post]
func (uC *UserController) Regist() {
	var newUser models.User
	json.Unmarshal(uC.Ctx.Input.RequestBody, &newUser)
	uC.Data["json"] = models.SetUser(&newUser)
	uC.ServeJSON()
}
