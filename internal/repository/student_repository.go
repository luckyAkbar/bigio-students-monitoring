package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) models.StudentRepository {
	return &studentRepo{
		db: db,
	}
}

func (r *studentRepo) FindByID(ctx context.Context, id int64) (*models.Student, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	student := &models.Student{}
	err := r.db.WithContext(ctx).Model(&models.Student{}).Where("id = ?", id).Take(&student).Error

	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return student, nil
	}
}

func (r *studentRepo) GetGradeBySubjectID(ctx context.Context, subjectID int64) (*models.Grade, error) {
	return nil, nil
}
