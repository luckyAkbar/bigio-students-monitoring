package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"user": utils.Dump(user),
	})

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}