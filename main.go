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
	//db.InitialMigration()
	log.Printf("Server started on localhost:%s\n", port)
	log.Print("Press Ctrl+C to Stop it")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
func main() {

	createServer()
}
