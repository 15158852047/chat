package main

import (
	_ "chat/routers"
	"github.com/astaxie/beego"
	_ "chat/initial"
)


func main() {
	beego.Run()
}

