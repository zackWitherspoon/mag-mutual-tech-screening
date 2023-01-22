package config

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"mag_matual_tech_project/constants"
	"net/http"
)

type Routes struct {
	Configuration *Configuration
}

func NewRoutes(configuration *Configuration) *Routes {
	return &Routes{
		Configuration: configuration,
	}
}
func (routes *Routes) LoadRoutes() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Endpoint to return a specific user (and all associated information)
	myRouter.HandleFunc(constants.GetUsersByIdEndpoint, routes.Configuration.Service.GetUserById)

	// Endpoint to return a specific user (and all associated information)
	myRouter.HandleFunc(constants.GetAllUsersEndpoint, routes.Configuration.Service.GetAllUsers)

	//// Endpoint to return a list of user created between a date range
	myRouter.HandleFunc(constants.GetUserInDateRangeEndpoint, routes.Configuration.Service.GetUsersByDatePage)

	// Endpoint to return a list of user based on a specific profession
	myRouter.HandleFunc(constants.GetUsersBasedOnProfessionEndpoint, routes.Configuration.Service.GetUsersByProfessionPage).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
