package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"web/models"
)

type UserloginController struct {
	beego.Controller
}

func (this *UserloginController) Get() {

	this.TplNames = "userlogin.html"
	return
}
func (this *UserloginController) Index() {

}
func (this *UserloginController) Post() {
	username := this.Input().Get("username")
	pws := this.Input().Get("pws")
	var (
		code    int
		message string
		id      string
	)
	code, message, id = models.LoginUser(username, pws)
	switch code {
	case 205:
		fmt.Printf("%s\n", message)
	case 200:
		s := models.Redis_session.Session(this.Ctx.ResponseWriter, this.Ctx.Request)
		s.Set("UserID", id)
		fmt.Printf("%s\n", message)
		this.Redirect("/login/welcomePage", 302)

	case 401:
		fmt.Printf("%s\n", message)
	}
	return

}
func (this *UserloginController) WelcomePage() {
	this.TplNames = "welcomePage.html"
	return
}
