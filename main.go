package main

import (
	"handlers"
	"log"
	"net/http"
	"proxy"

	yaml "gopkg.in/yaml.v2"

	"io/ioutil"

	"fmt"

	"github.com/gorilla/mux"
)

func main() {

	data, err := ioutil.ReadFile("./proxy.yaml")
	fmt.Println(string(data))
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	obj := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &obj)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	fmt.Printf("--- m:\n%v\n\n", obj["proxies"])

	p := proxy.New("http://mocktarget.apigee.net")

	r := mux.NewRouter()

	for _, proxyData := range obj["proxies"].(map[interface{}]interface{}) {

		m2 := make(map[string]string)
		for key, value := range proxyData.(map[interface{}]interface{}) {
			switch key := key.(type) {
			case string:
				switch value := value.(type) {
				case string:
					m2[key] = value
				}
			}
		}

		p = proxy.New(m2["url"])
		r.Handle(m2["base_path"], http.StripPrefix(m2["base_path"], p))
	}

	r.HandleFunc("/rand", handlers.HandleGenerateSig)
	r.HandleFunc("/foo", handlers.HandleFooRequest)
	r.HandleFunc("/keys", handlers.HandleGenerateKeyPair)
	r.HandleFunc("/storage", handlers.HandleStoreKeyValue).Methods("POST")
	r.HandleFunc("/storage/{key}", handlers.HandleGetKeyValue)
	//r.HandleFunc("/", handlers.HandleBasicRequest)

	r.HandleFunc("/keys/{id}", handlers.HandlePostGenerateKey).Methods("POST")

	r.Handle("/proxy", p)

	log.Fatal(http.ListenAndServe(":8080", r))
}
