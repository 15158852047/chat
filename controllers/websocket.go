package controllers

import (
	"chat/initial"
	"chat/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"
	"time"
)

type WebSocketControll struct {
	beego.Controller
}

var Active = make(map[*websocket.Conn]string)
var Upgrade = websocket.Upgrader{}
var Hellochan = make(chan models.HelloMsg)
var Chatchan = make(chan models.Chat)

func (this *WebSocketControll) Get() {
	ws, err := Upgrade.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)

	if err != nil {
		this.ServeJSON()
		return
	}
	id := this.GetSession("id").(int)
	username := this.GetSession("name").(string)
	Active[ws] = username
	this.SetSession("ws", ws)

	go ReadMsg(ws, this)

	o := orm.NewOrm()
	var user models.Users

	o.Raw("SELECT * FROM users WHERE id = ?", id).QueryRow(&user)

	var msg = models.HelloMsg{Name: user.Name, Time: initial.GetNowStr() + "好！", Username: user.Username, Email: user.Email, Ws: ws}
	Hellochan <- msg
}

func ReadMsg(ws *websocket.Conn, this *WebSocketControll) {
	m := models.Recive{}
	for {
		err := ws.ReadJSON(&m)
		if err == websocket.ErrCloseSent {
			break
		} else if err != nil {
			break
		} else {
			if m.Code != "" {
				if m.Code == "1" {
					from := this.GetSession("name").(string)
					i := this.GetSession("id").(int)
					to := m.To
					msg := m.Msg
					times := time.Now()

					o := orm.NewOrm()
					o.Raw("INSERT INTO messages (froms,tos,time,info,is_read) VALUES (?,?,?,?,?)", from, to, times, msg, 0).Exec()

					var toid int
					o.Raw("SELECT id FROM users WHERE username = ?", to).QueryRow(&toid)

					var id int
					o.Raw("SELECT id FROM chat_list where users_id = ? and username = ?", i, to).QueryRow(&id)
					if id != 0 {
						o.Raw("UPDATE chat_list SET info=?,time=? WHERE users_id = ? and username = ?", msg, times, i, to).Exec()
						o.Raw("UPDATE chat_list SET info=?,time=? WHERE users_id = ? and username = ?", msg, times, toid, from).Exec()
					} else {
						o.Raw("INSERT INTO chat_list (username,info,time,users_id) VALUES(?,?,?,?)", to, msg, times, i).Exec()
						o.Raw("INSERT INTO chat_list (username,info,time,users_id) VALUES(?,?,?,?)", from, msg, times, toid).Exec()
					}

					wss := make([]*websocket.Conn, 0)
					for key, value := range Active {
						if value == to || value == from {
							wss = append(wss, key)
						}
					}

					onemsg := models.Chat{From: from, To: to, Msg: msg, Time: times, Wss: wss}
					Chatchan <- onemsg
				}
			} else if m.Code == "100" {
				fmt.Println("100")
				ws.Close()

			}
		}
	}
}

func init() {
	go func() {
		for {
			select {
			case msg := <-Hellochan:
				{
					var data = map[string]interface{}{"code": 1, "name": msg.Name, "time": msg.Time, "username": msg.Username, "email": msg.Email}
					client := msg.Ws
					err := client.WriteJSON(data)
					if err != nil {
						client.Close()
						delete(Active, client)
					}
				}
			case msg := <-Chatchan:
				{
					var data = map[string]interface{}{"code": 2, "from": msg.From, "to": msg.To, "msg": msg.Msg, "time": msg.Time}
					clients := msg.Wss
					for _, c := range clients {
						err := c.WriteJSON(data)
						if err != nil {
							c.Close()
							delete(Active, c)
						}
					}
				}

			}
		}
	}()
}
