package models

import "time"

type Messages struct {
	Id     int
	Froms  string
	Tos    string
	Time   time.Time `orm:"auto_now_add;type(datetime)"`
	Info   string
	IsRead int //1 read  0 no
}
