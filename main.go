package main

import (
	"fmt"
	"log"
	"net/http"
)

func routes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func handleRequest() {
	http.HandleFunc("/routes", routes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequest()
}
