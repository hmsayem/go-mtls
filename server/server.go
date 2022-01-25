package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Println(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
		Handler:   router,
	}
	log.Fatalln(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
