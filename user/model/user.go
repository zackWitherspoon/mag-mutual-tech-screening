package model

import (
	"fmt"
	"time"
)

type User struct {
	id          int
	firstName   string
	lastName    string
	email       string
	profession  string
	dateCreated time.Time
	country     string
	city        string
}

func (user *User) Print() {
	fmt.Println(user.id, user.firstName, user.lastName, user.email, user.email,
		user.profession, user.dateCreated.String(), user.country, user.city)
}
