package routers

import (
	. "chat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &MainController{})

	beego.Router("/regist", &RegistController{})
	beego.Router("/login", &LoginController{})
	beego.Router("/logout", &LoginController{})

	beego.Router("/ws", &WebSocketControll{})
	beego.Router("/friendlist", &FriendController{})
	beego.Router("/chatlist", &ChatListController{})
	beego.Router("/chatMsg", &ChatMsgController{})

	beego.Router("/delmsg", &DelMsgController{})
}
