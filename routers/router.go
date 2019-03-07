package routers

import (
	. "chat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &MainController{})

    beego.Router("/regist",&RegistController{})
    beego.Router("/login",&LoginController{})
    beego.Router("/logout",&LoginController{})

    beego.Router("/ws",&WebSocketControll{})
}
