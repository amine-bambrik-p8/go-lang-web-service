package main

import (
	"go-lang-web-service/api"
	"log"
	"net/http"
)

/**
 */

func createServer() {
	api.HookRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	createServer()
}
