package db

import (
	"forum/internal/models"
	"time"

	"github.com/google/uuid"
)

func CreateCookie(userID int) (string, error) {
	token := uuid.NewString()
	expires := time.Now().Add(24 * time.Hour)

	_, err := DB.Exec(
		"INSERT INTO sessions (user_id, token, expires_at) VALUES (?, ?, ?)",
		userID, token, expires,
	)
	return token, err

}

func GetCookie(token string) (*models.Session, error) {
	s := &models.Session{}
	err := DB.QueryRow(
		"SELECT id, user_id, token, expires_at FROM sessions WHERE token = ? AND expires_at > ?",
		token, time.Now(),
	).Scan(&s.ID, &s.UserID, &s.Token, &s.ExpiresAt)
	return s, err
}
