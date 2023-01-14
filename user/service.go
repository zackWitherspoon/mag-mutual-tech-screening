package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mag_matual_tech_project/constants"
	"mag_matual_tech_project/user/model"
	"net/http"
	"strings"
	"time"
)

type V1Service struct {
	Users *[]model.User
}

func (service *V1Service) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetAllUsers")
	response, _ := json.MarshalIndent(service.Users, "", "    ")
	fmt.Fprintf(w, "%v", string(response))

}

func (service *V1Service) GetUsersByDatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUsersByDate")

	query := r.URL.Query()
	fromDate, _ := time.Parse(constants.DateFormat, query.Get("fromDate"))
	toDate, _ := time.Parse(constants.DateFormat, query.Get("toDate"))
	var foundUsers []model.User
	for _, user := range *service.Users {
		if user.DateCreated.After(fromDate) && user.DateCreated.Before(toDate) || user.DateCreated.Equal(fromDate) ||
			user.DateCreated.Equal(toDate) {
			foundUsers = append(foundUsers, user)
		}
	}

	response, _ := json.MarshalIndent(foundUsers, "", "    ")
	fmt.Fprintf(w, "%v", string(response))
}

func (service *V1Service) GetUsersByProfessionPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUsersByProfession")
	body, _ := ioutil.ReadAll(r.Body)
	print(string(body))
	var reqProfession model.ProfessionRequest
	if err := json.Unmarshal(body, &reqProfession); err != nil {
		// TODO: Log errors here
	}

	var foundUsers []model.User
	for _, user := range *service.Users {
		if strings.ToLower(user.Profession) == strings.ToLower(reqProfession.Profession) {
			foundUsers = append(foundUsers, user)
		}
	}
	response, _ := json.MarshalIndent(foundUsers, "", "    ")
	fmt.Fprintf(w, "%v", string(response))
}
