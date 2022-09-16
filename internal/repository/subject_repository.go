package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type subjectRepo struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) models.SubjectRepository {
	return &subjectRepo{
		db: db,
	}
}

func (r *subjectRepo) FindByID(ctx context.Context, id int64) (*models.Subject, error) {
	return nil, nil
}

func (r *subjectRepo) Create(ctx context.Context, subject *models.Subject) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":     utils.DumpIncomingContext(ctx),
		"subject": utils.Dump(subject),
	})

	if err := r.db.WithContext(ctx).Create(subject).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
