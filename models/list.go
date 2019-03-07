package models

type List struct {
	Id int
	Username string
	Users *Users `orm:"rel(fk)"`
}
