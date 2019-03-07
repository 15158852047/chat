package controllers

import (
	"chat/initial"
	"chat/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"
)

type WebSocketControll struct {
	beego.Controller
}

var Active = make(map[*websocket.Conn]int)
var Upgrade = websocket.Upgrader{}
var Hellochan = make(chan models.HelloMsg)

func (this *WebSocketControll) Get() {
	ws ,err := Upgrade.Upgrade(this.Ctx.ResponseWriter,this.Ctx.Request,nil)

	if err != nil {
		this.ServeJSON()
		return
	}

	id := this.GetSession("id").(int)
	Active[ws] = id
	this.SetSession("ws",ws)

	o := orm.NewOrm()
	var str string
	o.Raw("SELECT name FROM users WHERE id = ?",id).QueryRow(&str)

	var msg = models.HelloMsg{Name:str,Time:initial.GetNowStr()+"好！",Ws:ws}
	Hellochan <- msg
}


func init() {
	go func() {
		for{
			select {
			case msg := <- Hellochan:{
				var data = map[string]interface{}{"code":1,"name":msg.Name,"time":msg.Time}
				client := msg.Ws
				err := client.WriteJSON(data)
				if err != nil {
					client.Close()
					delete(Active,client)
				}
			}
			}
		}
	}()
}