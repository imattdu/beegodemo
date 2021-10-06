package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id       int
	Username string
	Age      int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
