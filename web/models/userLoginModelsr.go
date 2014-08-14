package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Customers struct {
	Id       int64
	Username string `orm:"index no null"`
	Password string
}

func LoginUser(name, pws string) (code int, message string, id string) {
	o := orm.NewOrm()
	seekUser := &Customers{Username: name}
	qs := o.QueryTable("customers")
	err := qs.Filter("username", name).One(seekUser)
	if err != nil {
		fmt.Println(err)
		code, message = 205, "您还没有注册吧！"
		return
	} else if pws == seekUser.Password {
		fmt.Printf("%T = %v\n", seekUser.Id, seekUser.Id)

		id = strconv.FormatInt(seekUser.Id, 10)
		fmt.Printf("%T = %v\n", id, id)
		code, message = 200, "登录成功"
		return
	} else {
		code, message = 401, "密码不正确"
		return
	}
	return

}
