package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohadese-yousefi/gymbell/internal/config"
	"github.com/mohadese-yousefi/gymbell/internal/models"
	"github.com/mohadese-yousefi/gymbell/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(config.GetEnv("JWTSECRET", "mySecretKey"))
type AuthService struct {
    Repo *repositories.UserRepository
}

func (s *AuthService) Register(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.Repo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
    user, err := s.Repo.GetUserByEmail(email)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // 1-day expiration
    })

    return token.SignedString(jwtSecret)
}
