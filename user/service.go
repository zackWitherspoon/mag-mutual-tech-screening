package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mag_matual_tech_project/constants"
	"mag_matual_tech_project/user/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	idErrorMessage                   = "Id used in an incorrect format. Try again with a different id"
	dateUnParsableErrorMessage       = "fromDate or toDate unable to be parsed. Please try again with a different date"
	professionUnParsableErrorMessage = "Unable to unmarshal request. Try with a different profession"
)

type V1Service struct {
	Users *[]model.User
}

func (service *V1Service) GetAllUsers(w http.ResponseWriter, request *http.Request) {
	log.Info("Endpoint Hit: GetAllUsers")

	apiResponse := model.APIResponse{Users: *service.Users}
	response, err := json.MarshalIndent(apiResponse, "", "    ")
	if err != nil {
		log.Error("Unable to marshal apiResponse: ", apiResponse, " into response object.")
	}
	log.Error("Returning result: ", apiResponse)
	fmt.Fprintf(w, "%v", string(response))

}

func (service *V1Service) GetUserById(writer http.ResponseWriter, request *http.Request) {
	log.Info("Endpoint Hit: etUsersById")
	requestVariables := mux.Vars(request)
	id, err := strconv.Atoi(requestVariables["id"])
	if err != nil {
		error := fmt.Sprint(idErrorMessage)
		log.Error(error, err)
		http.Error(writer, error, http.StatusBadRequest)
		return
	}

	var foundUsers []model.User
	for _, user := range *service.Users {
		if user.Id == id {
			foundUsers = append(foundUsers, user)
		}
	}
	apiResponse := model.APIResponse{Users: foundUsers}
	response, err := json.MarshalIndent(apiResponse, "", "    ")
	if err != nil {
		log.Error("Unable to marshal apiResponse: ", apiResponse, " into response object.")
	}

	log.Error("Returning result: ", apiResponse)
	fmt.Fprintf(writer, "%v", string(response))
}

func (service *V1Service) GetUsersByDatePage(writer http.ResponseWriter, request *http.Request) {
	log.Info("Endpoint Hit: getUsersByDate")

	query := request.URL.Query()
	fromDate, err := time.Parse(constants.DateFormat, query.Get("fromDate"))
	toDate, err := time.Parse(constants.DateFormat, query.Get("toDate"))
	if err != nil {
		log.Error(dateUnParsableErrorMessage, err)
		http.Error(writer, dateUnParsableErrorMessage, http.StatusBadRequest)
		return
	}
	var foundUsers []model.User
	for _, user := range *service.Users {
		if user.DateCreated.After(fromDate) && user.DateCreated.Before(toDate) || user.DateCreated.Equal(fromDate) ||
			user.DateCreated.Equal(toDate) {
			foundUsers = append(foundUsers, user)
		}
	}
	apiResponse := model.APIResponse{Users: foundUsers}
	response, err := json.MarshalIndent(apiResponse, "", "    ")
	if err != nil {
		log.Error("Unable to marshal apiResponse: ", apiResponse, " into response object.")
	}

	log.Error("Returning result: ", apiResponse)
	fmt.Fprintf(writer, "%v", string(response))
}

func (service *V1Service) GetUsersByProfessionPage(writer http.ResponseWriter, request *http.Request) {
	log.Info("Endpoint Hit: getUsersByProfession")
	body, _ := ioutil.ReadAll(request.Body)
	log.Infof(string(body))
	var reqProfession model.APIProfessionRequest
	if err := json.Unmarshal(body, &reqProfession); err != nil {
		log.Error(professionUnParsableErrorMessage, err)
		http.Error(writer, professionUnParsableErrorMessage, http.StatusBadRequest)
		return
	}

	var foundUsers []model.User
	for _, user := range *service.Users {
		if strings.ToLower(user.Profession) == strings.ToLower(reqProfession.Profession) {
			foundUsers = append(foundUsers, user)
		}
	}

	apiResponse := model.APIResponse{Users: foundUsers}
	response, err := json.MarshalIndent(apiResponse, "", "    ")
	if err != nil {
		log.Error("Unable to marshal apiResponse: ", apiResponse, " into response object.")
	}

	log.Error("Returning result: ", apiResponse)
	fmt.Fprintf(writer, "%v", string(response))
}
