package integration

import "mag_matual_tech_project/user/model"

type csvUsers struct {
}

func (csv *csvUsers) GetUsersByColumn(criteria string, users *[]model.User) *[]model.User {
	return &[]model.User{}
}
