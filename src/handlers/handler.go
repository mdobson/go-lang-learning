package handlers

import (
	"fmt"
	"hello"
	"io/ioutil"
	"net/http"
	"persist"
	"strings"

	"github.com/gorilla/mux"
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

func HandleGenerateKeyPair(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hello.GenerateKeys())
}

func HandlePostGenerateKey(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("foo: %s", r.PostForm)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ERROR: %s\n", err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "ECHO: %s\n", string(body))
	}
}

func HandleStoreKeyValue(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	keyValue := strings.Split(string(body), "=")
	persist.Save(keyValue[0], keyValue[1])

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK: %s\n", string(body))
}

func HandleGetKeyValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	keyValue, _ := persist.Get(key)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK: %s\n", keyValue)
}
