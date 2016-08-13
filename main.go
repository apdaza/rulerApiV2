package main

import (
	_ "github.com/apdaza/rulerApiV2/docs"
	_ "github.com/apdaza/rulerApiV2/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://oasuser:oasuser@127.0.0.1/udistrital?sslmode=disable&search_path=ruler")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
