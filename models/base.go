// Package models provides ...
package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	var mysqlUser = beego.AppConfig.String("mysql::user")
	var mysqlPassword = beego.AppConfig.String("mysql::pass")
	var mysqlUrl = beego.AppConfig.String("mysql::url")
	var mysqlDbname = beego.AppConfig.String("mysql::dbname")
	var mysqlParams = beego.AppConfig.String("mysql::params")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
			mysqlUser,
			mysqlPassword,
			mysqlUrl,
			mysqlDbname,
			mysqlParams,
		),
	)
	orm.RegisterModel(new(User), new(Profile))
	orm.RunSyncdb("default", false, true)
}
