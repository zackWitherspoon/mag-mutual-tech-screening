package user

import (
	"mag_matual_tech_project/user/model"
	"net/http"
)

type Service interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUsersByDatePage(w http.ResponseWriter, r *http.Request)
	GetUsersByProfessionPage(w http.ResponseWriter, r *http.Request)
}

func NewService(users *[]model.User) Service {
	return &V1Service{
		Users: users,
	}
}
