package main

import (
	_ "confignodoapi/docs"
	_ "confignodoapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	"fmt"

)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:123456@127.0.0.1:5432/nodocentral")

	//sincroniza base de datos
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
	    fmt.Println(err)
	}

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

