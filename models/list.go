package models

type List struct {
	Id       int
	Username string
	Firstp   string
	Users    *Users `orm:"rel(fk)"`
}
