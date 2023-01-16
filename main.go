package main

import log "github.com/sirupsen/logrus"

func main() {
	log.Info("Application is starting up")
	routes := NewRoutes(Configuration{}.NewConfiguration())
	log.Info("Routes have been configured")
	routes.LoadRoutes()
}
