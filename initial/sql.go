package initial

import (
	"chat/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initsql()
}

func initsql () {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.Debug = true
	orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/chat?charset=utf8")

	orm.RegisterModel(new(models.Users),new(models.List),new(models.Messages))
	//orm.RunSyncdb("default",false,true)
}