package db

import (
	"errors"
	"forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

var ErrEmailTaken = errors.New("email already taken")

func IsEmailTaken(email string) (bool, error) {
	var count int
	if err := DB.QueryRow(
		"SELECT COUNT(*) FROM users WHERE email = ?",
		email,
	).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateUser(u models.UserInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		u.Username, u.Email, string(hash),
	)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := DB.QueryRow(
		"SELECT id, username, email, password, created_at FROM users WHERE email = ?",
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
