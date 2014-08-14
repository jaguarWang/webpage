package routers

import (
	"web/controllers"
	"web/models"
	/*"fmt"*/
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strconv"
)

var FilterUser = func(ctx *context.Context) {
	s := models.Redis_session.Session(ctx.ResponseWriter, ctx.Request)
	sess_uid := s.Get("UserID")
	abc, _ := strconv.Atoi(sess_uid)
	/*fmt.Printf("%T =%v\n", sess_uid, sess_uid)
	fmt.Printf("%T =%v\n", web, web)*/

	if abc == 0 {

		ctx.Redirect(302, "/")
		return
	}
}

func init() {
	beego.InsertFilter("/login/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.UserloginController{})
	beego.Router("/login/welcomePage", &controllers.UserloginController{}, "get:WelcomePage")
}
