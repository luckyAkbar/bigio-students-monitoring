package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type gradeRepo struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) models.GradeRepository {
	return &gradeRepo{
		db,
	}
}

func (r *gradeRepo) Create(ctx context.Context, grade *models.Grade) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"grade": utils.Dump(grade),
	})

	if err := r.db.WithContext(ctx).Create(grade).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
