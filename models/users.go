package models

type Users struct {
	Id       int
	Username string
	Password string
	Name     string
	Email    string
	Firstp   string
}

type User struct {
	Username string
	Password string
	Name     string
	Email    string
}
