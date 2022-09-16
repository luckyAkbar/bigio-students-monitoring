package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
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
	logger := logrus.WithFields(logrus.Fields{
		"ctx":     utils.DumpIncomingContext(ctx),
		"session": utils.Dump(session),
	})

	if err := r.db.WithContext(ctx).Create(session).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *sessionRepo) FindByAccessToken(ctx context.Context, accessToken string) (*models.Session, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"token": accessToken,
	})

	session := &models.Session{}

	err := r.db.WithContext(ctx).Model(&models.Session{}).
		Where("access_token = ?", accessToken).Take(session).Error

	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return session, nil
	}
}
