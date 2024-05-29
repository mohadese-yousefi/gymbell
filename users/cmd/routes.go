package main

import "github.com/gorilla/mux"

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).Methods("GET")
	return r
}