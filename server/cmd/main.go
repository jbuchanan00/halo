package main

import (
	"halo/internal/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello world")
	r := router.New()

	port := ":8080"

	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Server failed")
	}
}
