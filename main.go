package main

import (
	"log"
	"net/http"

	"github.com/amine-bambrik-p8/go-lang-web-service/handlers"
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
