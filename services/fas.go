package services

import (
	"github.com/astaxie/beego"
)

type Fas struct {
}

func GetPermission(this Fas) string {
	env := beego.BConfig.RunMode
	fasHost := beego.AppConfig.String(env + "::fas")
	return fasHost
}