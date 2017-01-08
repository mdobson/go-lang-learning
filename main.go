package main

import (
	"handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rand", handlers.HandleGenerateSig)
	r.HandleFunc("/foo", handlers.HandleFooRequest)
	r.HandleFunc("/keys", handlers.HandleGenerateKeyPair)
	r.HandleFunc("/", handlers.HandleBasicRequest)

	r.HandleFunc("/keys/{id}", handlers.HandlePostGenerateKey).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
