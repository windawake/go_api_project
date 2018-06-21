package main

import (
	_ "apiproject/routers"

	"github.com/astaxie/beego"
)

func main() {
	// fmt.Printf("\n%+v\n\n", beego.BConfig)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
