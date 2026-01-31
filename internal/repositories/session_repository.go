package repositories

import (
	"rest-api-go-gin/internal/database"
	"rest-api-go-gin/internal/models"
)

type SessionRepository struct{}

func (r *SessionRepository) Create(session *models.Session) error {
	return database.DB.Create(session).Error
}

func (r *SessionRepository) FindByToken(token string) (*models.Session, error) {
	var session models.Session
	err := database.DB.Where("token = ? AND revoked = ?", token, false).First(&session).Error
	return &session, err
}

func (r *SessionRepository) RevokeByToken(token string) error {
	return database.DB.Model(&models.Session{}).Where("token = ?", token).Update("revoked", true).Error
}

func (r *SessionRepository) RevokeAllForUser(userID uint) error {
	return database.DB.Model(&models.Session{}).Where("user_id = ?", userID).Update("revoked", true).Error
}
