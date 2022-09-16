package usecase

import (
	"context"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/auth"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/repository"
	"github.com/sirupsen/logrus"
)

type teacherUsecase struct {
	studentRepo models.StudentRepository
	subjectRepo models.SubjectRepository
	gradeRepo   models.GradeRepository
}

func NewTeacherUsecase(studentRepo models.StudentRepository, subjectRepo models.SubjectRepository, gradeRepo models.GradeRepository) models.TeacherUsecase {
	return &teacherUsecase{
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
		gradeRepo:   gradeRepo,
	}
}

func (u *teacherUsecase) FindByID(ctx context.Context, id int64) (*models.Teacher, error) {
	return nil, nil
}

func (u *teacherUsecase) GradeByStudentID(ctx context.Context, input *models.CreateGradeInput) (*models.Grade, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"input": utils.Dump(input),
	})

	user := auth.GetUserFromCtx(ctx)
	if user.Role != models.RoleTeacher {
		return nil, ErrUnauthorized
	}

	if err := input.Validate(); err != nil {
		return nil, ErrValidation
	}

	_, err := u.studentRepo.FindByID(ctx, input.StudentID)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	subject, err := u.subjectRepo.FindByID(ctx, input.SubjectID)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	if subject.TeacherID != user.ID {
		return nil, ErrUnauthorized
	}

	grade := &models.Grade{
		ID:        utils.GenerateID(),
		StudentID: input.StudentID,
		TeacherID: user.ID,
		SubjectID: input.SubjectID,
		Mark:      input.Mark,
		Value:     input.Value,
	}

	if err := u.gradeRepo.Create(ctx, grade); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return grade, nil
}
