package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amine-bambrik-p8/go-lang-web-service/handlers"
)

func createServer() {
	handlers.HookHandlers()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
func main() {
	createServer()
}
