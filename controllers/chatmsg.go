package controllers

import (
	"chat/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ChatMsgController struct {
	beego.Controller
}

func (this *ChatMsgController) Get() {
	toname := this.GetString("username")
	iname := this.GetSession("name").(string)
	pagecount, _ := this.GetInt("pagecount")

	o := orm.NewOrm()
	var msgs []models.Messages
	low := (pagecount - 1) * 10
	high := pagecount * 10
	o.Raw("SELECT * FROM (SELECT * FROM messages WHERE (froms = ? AND tos = ?) OR (froms = ? AND tos = ?) ORDER BY time DESC LIMIT ?,?) AS a ORDER BY time ", iname, toname, toname, iname, low, high).QueryRows(&msgs)

	var count int
	o.Raw("SELECT COUNT(*) FROM messages WHERE (froms = ? AND tos = ?) OR (froms = ? AND tos = ?)", iname, toname, toname, iname).QueryRow(&count)

	mm := make(map[string]interface{})
	if count > (pagecount * 10) {
		mm["haveMore"] = 1
	} else {
		mm["haveMore"] = 0
	}
	msglist := make([]map[string]string, 0)
	for _, msg := range msgs {
		m := make(map[string]string)
		if msg.Froms == iname {
			m["people"] = "me"
		} else {
			m["people"] = "other"
		}
		m["msg"] = msg.Info
		msglist = append(msglist, m)
	}
	mm["infos"] = msglist

	o.Raw("UPDATE messages SET is_read = 1 WHERE froms = ? AND tos = ?", toname, iname).Exec()

	this.Data["json"] = mm
	this.ServeJSON()
}
