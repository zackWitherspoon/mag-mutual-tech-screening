package model

import (
	"encoding/json"
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

func MarshalArray(users []User) []byte {
	var marshaledArray []byte
	for _, user := range users {
		payloadBytes, _ := json.Marshal(user)
		marshaledArray = append(marshaledArray, payloadBytes...)
	}
	return marshaledArray
}
