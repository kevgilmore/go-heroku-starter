package main

import (
	"github.com/gorilla/mux"
	"go-starter/config"
	"go-starter/controller"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on port: "+port)
	router := mux.NewRouter()
	router.Use(config.AddJsonHeader)

	//ROUTES
	router.HandleFunc("/health", controller.HomeHandler).Methods("GET")
	log.Printf("Mapped [/health] for server status")

	http.ListenAndServe(":"+port, router)
}
