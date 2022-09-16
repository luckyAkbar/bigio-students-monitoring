package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) models.AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) FindByID(ctx context.Context, id int64) (*models.Admin, error) {
	return nil, nil
}

func (r *adminRepository) Create(ctx context.Context, admin *models.Admin) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"admin": utils.Dump(admin),
	})

	if err := r.db.WithContext(ctx).Create(admin).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *adminRepository) CreateTeacher(ctx context.Context, teacher *models.Teacher, user *models.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":     utils.DumpIncomingContext(ctx),
		"teacher": utils.Dump(teacher),
		"user":    utils.Dump(user),
	})

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(teacher).Error; err != nil {
			logger.Error(err)
			return err
		}

		if err := tx.WithContext(ctx).Create(user).Error; err != nil {
			logger.Error(err)
			return err
		}

		return nil
	})

	return err
}

func (r *adminRepository) CreateStudent(ctx context.Context, teacher *models.Student, user *models.User) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":  utils.DumpIncomingContext(ctx),
		"user": utils.Dump(user),
	})

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(teacher).Error; err != nil {
			logger.Error(err)
			return err
		}

		if err := tx.WithContext(ctx).Create(user).Error; err != nil {
			logger.Error(err)
			return err
		}

		return nil
	})

	return err
}
