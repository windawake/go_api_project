// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apiproject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	apins := beego.NewNamespace("/api",
		beego.NSNamespace("/financialReport",
			beego.NSRouter("/getReportTypePermission", &controllers.FasController{},"get:GetReportTypePermission"),
		),
	)

	beego.AddNamespace(apins)

	// beego.Router("/api/financialReport/getReportTypePermission",&controllers.MainController{},"get:GetReportTypePermission")

	// beego.Router("/test", &controllers.MainController{})
	
}
