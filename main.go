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
	r.HandleFunc("/", handlers.HandleBasicRequest)

	log.Fatal(http.ListenAndServe(":8080", r))
}
