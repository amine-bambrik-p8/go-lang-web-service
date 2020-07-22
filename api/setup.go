package api

import (
	routes "go-lang-web-service/api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

//Hooks up all the api's endpoints
func HookRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/routes", routes.Controller.GetRoutes).Methods("GET")
	http.Handle("/", router)
}
