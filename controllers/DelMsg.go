package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DelMsgController struct {
	beego.Controller
}

type usjson struct {
	Us string `json:"us"`
}

func (this *DelMsgController) Post() {
	var a usjson
	json.Unmarshal(this.Ctx.Input.RequestBody, &a)
	id := this.GetSession("id").(int)

	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM chat_list WHERE users_id = ? AND username = ?", id, a.Us).Exec()

	if err != nil {
		this.Data["json"] = map[string]string{"code": "0"}
		this.ServeJSON()
	}

	this.Data["json"] = map[string]string{"code": "1"}
	this.ServeJSON()
}
