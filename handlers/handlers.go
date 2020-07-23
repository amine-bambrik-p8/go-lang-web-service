package handlers

import (
	"net/http"
)

// Hooks up all the api's routes
func HookHandlers() {
	http.Handle("/", RoutesHandler.HookHandlers())
	http.Handle("/users", UserHandler.HookHandlers())
}
