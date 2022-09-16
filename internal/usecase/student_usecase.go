package usecase

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/auth"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/repository"
	"github.com/sirupsen/logrus"
)

type studentUsecase struct {
	studentRepo models.StudentRepository
}

func NewStudentUsecase(studentRepo models.StudentRepository) models.StudentUsecase {
	return &studentUsecase{
		studentRepo,
	}
}

func (u *studentUsecase) FindByID(ctx context.Context, id int64) (*models.Student, error) {
	return nil, nil
}

func (u *studentUsecase) GetGradeBySubjectID(ctx context.Context, subjectID int64) (*models.Grade, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":       utils.DumpIncomingContext(ctx),
		"subjectID": subjectID,
	})

	user := auth.GetUserFromCtx(ctx)
	if user.Role != models.RoleStudent {
		return nil, ErrUnauthorized
	}

	grade, err := u.studentRepo.FindByStudentAndSubjectID(ctx, user.ID, subjectID)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		return grade, nil
	}
}

func (u *studentUsecase) GetAllGrade(ctx context.Context) (grades []models.Grade, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
	})

	user := auth.GetUserFromCtx(ctx)
	if user.Role != models.RoleStudent {
		return grades, ErrUnauthorized
	}

	grades, err = u.studentRepo.GetAllGrade(ctx, user.ID)
	if err != nil {
		logger.Error(err)
		return grades, ErrInternal
	}

	return grades, nil
}
