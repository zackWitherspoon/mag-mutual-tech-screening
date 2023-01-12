package integration

import "mag_matual_tech_project/user/model"

type Integrator interface {
	GetUsersByColumn(criteria string) []model.User
}

func NewIntegrator(users []model.User) Integrator {
	return &csvUsers{
		Users: users,
	}
}
