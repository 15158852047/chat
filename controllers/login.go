package controllers

import (
	"chat/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

type One struct {
	Name string
	Pass string
}

func (this LoginController) Post() {
	var one = &One{}
	json.Unmarshal(this.Ctx.Input.RequestBody, one)

	o := orm.NewOrm()
	var i int
	o.Raw("SELECT id FROM users WHERE username = ? and password = ?", one.Name, one.Pass).QueryRow(&i)
	if i > 0 {
		this.SetSession("id", i)
		this.SetSession("name", one.Name)
		m := map[string]interface{}{"code": 0, "value": "登陆成功!"}
		this.Data["json"] = m
		this.ServeJSON()
		return
	}

	m := map[string]interface{}{"code": 1, "value": "账号或密码错误!"}
	this.Data["json"] = m
	this.ServeJSON()
	return

}

type RegistController struct {
	beego.Controller
}

func (this *RegistController) Post() {
	var user = &models.User{}
	o := orm.NewOrm()
	json.Unmarshal(this.Ctx.Input.RequestBody, user)

	var i int
	o.Raw("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).QueryRow(&i)
	if i > 0 {
		m := map[string]interface{}{"code": 1, "value": "账号已经存在!"}
		this.Data["json"] = m
		this.ServeJSON()
		return
	}

	_, err := o.Raw("INSERT INTO users (username,password,name,email) VALUES (?,?,?,?)", user.Username, user.Password, user.Name, user.Email).Exec()
	if err != nil {
		m := map[string]interface{}{"code": 1, "value": "注册账号失败!"}
		this.Data["json"] = m
		this.ServeJSON()
		return
	}

	str := "http://127.0.0.1:5000/deal/" + user.Username
	http.Get(str)
	m := map[string]interface{}{"code": 0, "value": "注册账号成功!"}
	this.Data["json"] = m
	this.ServeJSON()
	return
}

type LogoutController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.DelSession("id")
	this.Redirect("/", 302)
}
