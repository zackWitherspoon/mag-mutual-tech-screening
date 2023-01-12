package main

import (
	"fmt"
	"mag_matual_tech_project/user/integration"
	"mag_matual_tech_project/user/model"
)

func main() {
	fmt.Println("hello world")
	integration.NewIntegrator([]model.User{})
	HandleRequests()
}
