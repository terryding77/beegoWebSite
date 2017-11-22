package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int `orm:"auto;pk"`
	Name    string
	Profile *Profile `orm:"rel(one)"`
}

func GetAllUsers() string {
	var o = orm.NewOrm()
	var lists []orm.ParamsList
	var qs = o.QueryTable("user")
	num, err := qs.ValuesList(&lists)
	if err == nil {
		return "get all users" + string(num)
	} else {
		return "error on get users"
	}
}

func SetUser(newUser *User) string {
	newUser.Profile = new(Profile)
	var o = orm.NewOrm()
	beego.Informational(*newUser)
	SetProfile(newUser.Profile)
	id, err := o.Insert(newUser)
	if err == nil {
		return "ok" + string(id)
	} else {
		return err.Error()
	}
}
