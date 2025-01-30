package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("your-secret-key")

// JwtCustomClaims holds custom claims for JWT
type JwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT for a user
func GenerateToken(userID int) (string, error) {
	claims := &JwtCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "echo-jwt-register",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

