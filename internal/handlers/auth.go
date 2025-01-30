package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohadese-yousefi/gymbell/internal/models"
	"github.com/mohadese-yousefi/gymbell/internal/services"
)

type AuthHandler struct {
    Service *services.AuthService
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.Service.Register(&user); err != nil {
		fmt.Println(err)
        http.Error(w, "Could not create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    token, err := h.Service.Login(credentials.Email, credentials.Password)
    if err != nil {
		fmt.Println(err)
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
