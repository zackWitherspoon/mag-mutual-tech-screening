package integration

import "mag_matual_tech_project/user/model"

type Integrator interface {
	GetUsersByColumn(criteria string, users *[]model.User) *[]model.User
}

func NewIntegrator() Integrator {
	return &csvUsers{}
}
