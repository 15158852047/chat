package controllers

import (
	"chat/initial"
	"chat/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type FriendController struct {
	beego.Controller
}

func (this *FriendController) Get() {
	name := this.GetSession("name").(string)
	id := this.GetSession("id").(int)

	o := orm.NewOrm()
	var list []models.List
	o.Raw("SELECT * FROM list WHERE users_id = ? ORDER BY firstp", id).QueryRows(&list)

	i := 0
	fp := ""
	lastfp := ""
	m := make(map[string]string)
	uselist := make(map[string]interface{}) //{"A":[{"name":"","message":""},{"name":"","message":""}],"B":[{}]}
	pinyin := make([]map[string]string, 100)

	for _, l := range list {
		if i == 0 {
			fp = l.Firstp
			username := l.Username
			m = GetOneUser(username, name)
			pinyin = []map[string]string{m}
			lastfp = fp
			i++
			if len(list) == 1 {
				uselist[fp] = pinyin
			}
		} else {
			username := l.Username
			m = GetOneUser(username, name)
			if l.Firstp == fp {
				pinyin = append(pinyin, m)
			} else {
				fp = l.Firstp
				uselist[lastfp] = pinyin
				pinyin = []map[string]string{m}
				lastfp = fp
				uselist[lastfp] = pinyin
			}
		}
	}
	this.Data["json"] = uselist
	this.ServeJSON()

}

func GetOneUser(username string, name string) map[string]string {
	m := make(map[string]string)
	o := orm.NewOrm()

	var user models.Users
	o.Raw("SELECT * FROM users WHERE username = ?", username).QueryRow(&user)

	var message models.Messages
	o.Raw("SELECT * FROM messages WHERE (froms = ? AND tos = ?) OR (froms = ? AND tos = ?) ORDER BY time DESC limit 1", user, name, name, user).QueryRow(&message)

	m["name"] = user.Name
	m["username"] = username
	m["msg"] = message.Info
	m["time"] = initial.GetMsgTime(message.Time)
	i := initial.GenerateRangeNum(1, 12)
	m["img"] = "static/images/head/" + strconv.Itoa(i) + ".jpg"

	return m
}
