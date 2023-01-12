package integration

import "mag_matual_tech_project/user/model"

type csvUsers struct {
	Users []model.User
}

func (csv *csvUsers) GetUsersByColumn(criteria string) []model.User {
	return []model.User{}
}
