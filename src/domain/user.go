package domain

import "time"

type User struct {
	Name string
	Mail string
	Nick string
	Pass string
	Date *time.Time
}

func NewUser(name string, mail string, nick string, pass string) *User {
	var date = time.Now()
	var creado = User{Name: name, Mail: mail, Nick: nick, Pass: pass, Date: &date}
	return &creado
}
