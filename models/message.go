package models

import "time"

type Messages struct {
	Id int
	From int
	To int
	Time time.Time `orm:"auto_now_add;type(datetime)"`
	Info string
	IsRead int
}
