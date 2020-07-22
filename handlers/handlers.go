package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Hooks up all the api's routes
func HookHandlers() {
	router := mux.NewRouter()
	router.HandleFunc("/routes", RoutesHandler.GetRoutes).Methods("GET")
	http.Handle("/", router)
}
