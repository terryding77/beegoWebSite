package routers

import (
	"goWebTest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//beego.NSGet("/version", aaa),
		beego.NSInclude(&controllers.CommonController{}),
		//beego.Include(&commonController{}),
	)

	beego.AddNamespace(ns)
}
