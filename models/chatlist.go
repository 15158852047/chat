package models

import "time"

type ChatList struct {
	Id       int
	Username string
	Info     string
	Time     time.Time
	Users    *Users `orm:"rel(fk)"`
}
