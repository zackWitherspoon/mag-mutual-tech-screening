package model

import (
	"time"
)

type User struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Profession  string    `json:"profession"`
	DateCreated time.Time `json:"dateCreated"`
	Country     string    `json:"country"`
	City        string    `json:"email"`
}

//func Print(users *[]User) {
//	for user := range users {
//
//	}
//	fmt.Println(user.Id, user.FirstName, user.LastName, user.Email, user.Profession, user.DateCreated.String(),
//		user.Country, user.City)
//}
