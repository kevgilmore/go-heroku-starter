package main

import (
	"github.com/gorilla/mux"
	"go-starter/controller"
	"net/http"
)

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.HomeHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
