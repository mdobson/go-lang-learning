package handlers

import (
	"fmt"
	"hello"
	"net/http"
)

func HandleBasicRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello.Hello())
}
