package repository

import (
	"context"

	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"gorm.io/gorm"
)

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) models.SessionRepository {
	return &sessionRepo{
		db: db,
	}
}

func (r *sessionRepo) Create(ctx context.Context, session *models.Session) error {
	return nil
}

func (r *sessionRepo) FindByID(ctx context.Context, id int64) (*models.Session, error) {
	return nil, nil
}

func (r *sessionRepo) FindByAccessToken(ctx context.Context, accessToken string) (*models.Session, error) {
	return nil, nil
}
