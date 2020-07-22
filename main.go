package main

import (
	"go-lang-web-service/handlers"
	"log"
	"net/http"
)

/**
 */

func createServer() {
	handlers.HookHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	createServer()
}
