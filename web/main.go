package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web/models"
	_ "web/routers"
)

func init() {
	/*初始化数据库*/
	models.RegisterDB()

}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	/* 设置Session配置*/
	beego.SessionOn = false
	beego.Run()
}
