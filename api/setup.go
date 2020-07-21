package api

import (
	routes "go-lang-web-service/api/routes"
	"net/http"
)

//Hooks up all the api's endpoints
func HookRoutes() {
	http.HandleFunc("/routes", routes.GetRoutes)
}
