package controllers

import (
	"beegodemo/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ListUser() {

	models.Insert("matt", 15)

	mystruct := map[string]string{"name": "matt"}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}
