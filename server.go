package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var router = mux.NewRouter()

func matcher(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Hello World!")
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

func handleRequest() {
	router.HandleFunc("/home", matcher).Methods("GET")
}

func main() {
	handleRequest()
	log.Fatalln(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", router))
}
