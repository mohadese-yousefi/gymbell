package api

import (
	"encoding/json"
	"github.com/mohadese-yousefi/gymbell/internal/config"
	"github.com/mohadese-yousefi/gymbell/internal/middleware"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}


func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db := config.GetDB()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE email=$1", user.Email)
	if err != nil {
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Insert new user using prepared statement
	query := `INSERT INTO users (username, email, password) VALUES (:username, :email, :password)`
	result, err := db.NamedExec(query, user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Generate token
	newUserId, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
		return
	}

	token, err := middleware.GenerateToken(int(newUserId))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"token":   token,
	})
}