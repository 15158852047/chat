package controllers

import (
	"chat/initial"
	"chat/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ChatListController struct {
	beego.Controller
}

func (this *ChatListController) Get() {
	id := this.GetSession("id").(int)
	usname := this.GetSession("name").(string)
	o := orm.NewOrm()

	var lists []models.ChatList
	o.Raw("SELECT * FROM chat_list WHERE users_id = ? ORDER BY time", id).QueryRows(&lists)

	chatlist := make([]map[string]string, 0)
	var s string
	var notRead int
	for _, list := range lists {
		o.Raw("SELECT name FROM users WHERE username=?", list.Username).QueryRow(&s)
		o.Raw("SELECT COUNT(*) FROM (SELECT * FROM messages WHERE froms = ? AND tos = ?) AS a WHERE is_read = 0", list.Username, usname).QueryRow(&notRead)
		m := make(map[string]string)
		m["name"] = s
		m["username"] = list.Username
		m["time"] = initial.GetMsgTime(list.Time)
		m["lastMessage"] = list.Info
		i := initial.GenerateRangeNum(1, 12)
		m["img"] = "static/images/head/" + strconv.Itoa(i) + ".jpg"
		m["notRead"] = strconv.Itoa(notRead)
		fmt.Println(notRead)
		chatlist = append(chatlist, m)
	}
	this.Data["json"] = chatlist
	this.ServeJSON()
}
