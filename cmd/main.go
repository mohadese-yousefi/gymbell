package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	v1 "github.com/mohadese-yousefi/gymbell/api/v1"
	"github.com/mohadese-yousefi/gymbell/internal/config"
	"github.com/mohadese-yousefi/gymbell/internal/handlers"
	"github.com/mohadese-yousefi/gymbell/internal/repositories"
	"github.com/mohadese-yousefi/gymbell/internal/services"
)

func main() {
    db := config.GetDB()
    defer db.Close()

    userRepo := &repositories.UserRepository{DB: db}
    authService := &services.AuthService{Repo: userRepo}
    authHandler := &handlers.AuthHandler{Service: authService}

    mux := http.NewServeMux()
    v1.RegisterRoutes(mux, authHandler)

    log.Println("Server running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
