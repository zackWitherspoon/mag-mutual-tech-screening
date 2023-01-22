package main

import (
	log "github.com/sirupsen/logrus"
	"mag_matual_tech_project/config"
)

func main() {
	log.Info("Application is starting up")
	routes := config.NewRoutes(config.Configuration{}.NewConfiguration())
	log.Info("Routes have been configured")
	routes.LoadRoutes()
}
