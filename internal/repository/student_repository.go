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

func (r *studentRepo) FindByStudentAndSubjectID(ctx context.Context, studentID, subjectID int64) (*models.Grade, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       utils.DumpIncomingContext(ctx),
		"studentID": studentID,
		"subjectID": subjectID,
	})

	grade := &models.Grade{}

	err := r.db.WithContext(ctx).Model(&models.Grade{}).
		Where("student_id = ? AND subject_id = ?", studentID, subjectID).
		Take(grade).Error

	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return grade, nil
	}
}

func (r *studentRepo) GetAllGrade(ctx context.Context, id int64) (grades []models.Grade, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
	})

	err = r.db.WithContext(ctx).Model(&models.Grade{}).
		Where("student_id = ?", id).Scan(&grades).Error

	if err != nil {
		logger.Error(err)
		return grades, err
	}

	return grades, nil
}
