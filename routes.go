package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	getAllUsersEndpoint               = "/users"
	getUserInDateRangeEndpoint        = "/getAllUsersEndpoint/date/{date}"
	getUsersBasedOnProfessionEndpoint = "/getAllUsersEndpoint/profession/{profession}"
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
	myRouter.HandleFunc(getAllUsersEndpoint, routes.Configuration.Service.GetUsersByCriteriaPage)
	myRouter.HandleFunc("/", homePage)
	//// Endpoint to return a list of user created between a date range
	myRouter.HandleFunc(getUserInDateRangeEndpoint, homePage)
	//
	// Endpoint to return a list of user based on a specific profession
	myRouter.HandleFunc(getUsersBasedOnProfessionEndpoint, homePage)
	//
	//// Custom Endpoint that you design on your own
	//myRouter.HandleFunc("/user/add", homePage).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
