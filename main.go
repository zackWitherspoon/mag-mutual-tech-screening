package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	routes := NewRoutes(Configuration{}.NewConfiguration())
	routes.LoadRoutes()
}
