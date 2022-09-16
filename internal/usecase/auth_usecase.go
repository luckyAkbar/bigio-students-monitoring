package usecase

import (
	"context"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/config"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/helper"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/repository"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	sessionRepo models.SessionRepository
	userRepo    models.UserRepository
}

func NewAuthUsecase(sessionRepo models.SessionRepository, userRepo models.UserRepository) models.AuthUsecase {
	return &authUsecase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (u *authUsecase) LoginByIDAndPassword(ctx context.Context, id int64, password string) (*models.Session, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	user, err := u.userRepo.FindByID(ctx, id)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	if helper.IsHashedStringMatch([]byte(password), []byte(user.Password)) {
		logger.Info("Failed login process detected")
		return nil, ErrUnauthorized
	}

	session := &models.Session{
		ID:          utils.GenerateID(),
		AccessToken: helper.GenerateToken(user.ID),
		UserID:      user.ID,
		Role:        user.Role,
		CreatedAt:   time.Now(),
		ExpiredAt:   time.Now().Add(config.DefaultAccessTokenExpiry),
	}

	if err := u.sessionRepo.Create(ctx, session); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return session, nil
}
