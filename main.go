package main

func main() {
	routes := NewRoutes(Configuration{}.NewConfiguration())
	routes.LoadRoutes()
}
