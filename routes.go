package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Endpoint to return a specific user (and all associated information)
	//myRouter.HandleFunc("/user", GetUsersByColumn)

	// Endpoint to return a list of user created between a date range
	myRouter.HandleFunc("/users/date/{date}", homePage)

	// Endpoint to return a list of user based on a specific profession
	myRouter.HandleFunc("/users/profession/{profession}", homePage)

	// Custom Endpoint that you design on your own
	myRouter.HandleFunc("/user/add", homePage).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
