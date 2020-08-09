package main

import (
	"go-service/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Dialog Service called!")
	handleRequest()
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/api/dialogs", api.Handle).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}