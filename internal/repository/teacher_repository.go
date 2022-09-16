package repository

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type teacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) models.TeacherRepository {
	return &teacherRepo{
		db: db,
	}
}

func (r *teacherRepo) FindByID(ctx context.Context, id int64) (*models.Teacher, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	teacher := &models.Teacher{}
	err := r.db.WithContext(ctx).Model(&models.Teacher{}).
		Where("id = ?", id).Take(teacher).Error

	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return teacher, nil
	}
}

func (r *teacherRepo) GradeByStudentID(ctx context.Context, grade *models.Grade) error {
	return nil
}
