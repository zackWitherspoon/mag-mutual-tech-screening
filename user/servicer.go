package user

import (
	"mag_matual_tech_project/user/integration"
	"mag_matual_tech_project/user/model"
	"net/http"
)

type Service interface {
	GetUsersByCriteriaPage(w http.ResponseWriter, r *http.Request)
}

func NewService(users *[]model.User, integrator integration.Integrator) Service {
	return &V1Service{
		Users:      users,
		integrator: integrator,
	}
}
