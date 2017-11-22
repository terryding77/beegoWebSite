// Package models provides ...
package models

import (
	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id   int `orm:"auto;pk"`
	Age  int16
	User *User `orm:"reverse(one)"`
}

func SetProfile(profile *Profile) string {
	var o = orm.NewOrm()
	id, err := o.Insert(profile)
	if err == nil {
		return "ok" + string(id)
	} else {
		return "error"
	}
}
