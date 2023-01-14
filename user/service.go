package user

import (
	"fmt"
	"mag_matual_tech_project/user/integration"
	"mag_matual_tech_project/user/model"
	"net/http"
)

type V1Service struct {
	Users      *[]model.User
	integrator integration.Integrator
}

func (service *V1Service) GetUsersByCriteriaPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", *service.Users)
	fmt.Println("Endpoint Hit: USERS HERE")
}
