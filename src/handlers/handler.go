package handlers

import (
	"fmt"
	"hello"
	"net/http"
)

func HandleBasicRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello.Hello())
}

func HandleFooRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello.Foo())
}

func HandleGenerateSig(w http.ResponseWriter, r *http.Request) {
	messages := make(chan string)

	go func() {
		messages <- hello.RandToken()
	}()

	msg := <-messages
	fmt.Fprint(w, msg)
}
