package routers

import (
	"github.com/astaxie/beego"
	"goWebTest/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSInclude(&controllers.CommonController{}),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
	)

	beego.AddNamespace(ns)
}
