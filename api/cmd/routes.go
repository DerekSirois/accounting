package main

import (
	"accounting/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.VerifyJWT(handlers.Index)).Methods("GET")

	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/register", handlers.Register).Methods("POST")

	return router
}
