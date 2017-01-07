package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/rand", handlers.HandleGenerateSig)
	http.HandleFunc("/foo", handlers.HandleFooRequest)
	http.HandleFunc("/", handlers.HandleBasicRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
