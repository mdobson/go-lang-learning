package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HandleBasicRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
