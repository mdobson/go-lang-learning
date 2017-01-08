package main

import (
	"handlers"
	"log"
	"net/http"
	"proxy"

	"github.com/gorilla/mux"
)

func main() {

	p := proxy.New("http://mocktarget.apigee.net")

	r := mux.NewRouter()
	r.HandleFunc("/rand", handlers.HandleGenerateSig)
	r.HandleFunc("/foo", handlers.HandleFooRequest)
	r.HandleFunc("/keys", handlers.HandleGenerateKeyPair)
	r.HandleFunc("/", handlers.HandleBasicRequest)

	r.HandleFunc("/keys/{id}", handlers.HandlePostGenerateKey).Methods("POST")

	r.HandleFunc("/proxy", p.Handle)

	log.Fatal(http.ListenAndServe(":8080", r))
}
