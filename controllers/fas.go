package controllers

import (
	"apiproject/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	// "fmt"
)

type FasController struct {
	beego.Controller
}

func (this *FasController) GetReportTypePermission() {
	// 获取请求体传参
	key := this.Input().Get("key")
	parentId := this.Input().Get("parentId")
	token := this.Ctx.Input.Header("token")

	// 获取远程服务地址
	var sFas services.Fas
	fasHost := services.GetPermission(sFas,);
	url := fasHost+"/api/financialReport/getReportTypePermission?key="+key+"&parentId="+parentId


	// 远程调用
	logs.SetLogger(logs.AdapterFile, `{"filename":"storage/web.log"}`)
	req := httplib.Get(url)
	req.Header("token", token)
	// fmt.Printf("\n%+v\n\n", req)
	// logs.Info("请求记录")

	str, err := req.String()
	if err != nil {
		logs.Error(err)
	}
  this.Ctx.Output.Body([]byte(str))
}
