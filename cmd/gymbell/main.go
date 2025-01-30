package main

import (
	"log"
	"net/http"
	"github.com/mohadese-yousefi/gymbell/internal/config"
	"github.com/mohadese-yousefi/gymbell/internal/api"
)

func main() {
	db := config.GetDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/register", api.Register)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}