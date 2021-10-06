package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func Insert(username string, age int) {
	o := orm.NewOrm()
	var user User
	user.Username = username
	user.Age = age

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}

}
