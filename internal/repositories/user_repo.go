package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/mohadese-yousefi/gymbell/internal/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (repo *UserRepository) CreateUser(user *models.User) error {
    query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	return repo.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    user := &models.User{}
    query := "SELECT id, username, email, password FROM users WHERE email = $1"
    err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    if err != nil {
        return nil, err
    }
    return user, nil
}
