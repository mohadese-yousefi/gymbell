package v1

import (
	"net/http"

	"github.com/mohadese-yousefi/gymbell/internal/handlers"
	"github.com/mohadese-yousefi/gymbell/pkg/middleware"
)

func RegisterRoutes(mux *http.ServeMux, authHandler *handlers.AuthHandler) {
    mux.HandleFunc("POST /api/v1/register", authHandler.Register)
    mux.HandleFunc("POST /api/v1/login", authHandler.Login)

    // Protected Route Example
    mux.Handle("/api/v1/protected", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("You have accessed a protected route"))
    })))
}
