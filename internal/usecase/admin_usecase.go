package usecase

import (
	"context"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/auth"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/repository"
	"github.com/sirupsen/logrus"
)

type adminUsecase struct {
	adminRepo   models.AdminRepository
	teacherRepo models.TeacherRepository
	subjectRepo models.SubjectRepository
}

func NewAdminUsecase(adminRepo models.AdminRepository, teacherRepo models.TeacherRepository, subjectRepo models.SubjectRepository) models.AdminUsecase {
	return &adminUsecase{
		adminRepo:   adminRepo,
		teacherRepo: teacherRepo,
		subjectRepo: subjectRepo,
	}
}

func (u *adminUsecase) FindByID(ctx context.Context, id int64) (*models.Admin, error) {
	return nil, nil
}

func (u *adminUsecase) CreateTeacher(ctx context.Context, input *models.CreateTeacherInput) (*models.Teacher, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"input": utils.Dump(input),
	})

	if err := input.Validate(); err != nil {
		return nil, ErrValidation
	}

	user := auth.GetUserFromCtx(ctx)
	if user.Role != models.RoleAdmin {
		logger.Info("unauthorized access on create teacher by: ", user)
		return nil, ErrUnauthorized
	}

	teacherID := utils.GenerateID()
	teacher := &models.Teacher{
		ID:   teacherID,
		Name: input.Name,
	}

	newUser := &models.User{
		ID:        teacherID,
		Password:  input.Password,
		CreatedAt: time.Now(),
		Role:      models.RoleTeacher,
	}

	if err := newUser.Encrypt(); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	if err := u.adminRepo.CreateTeacher(ctx, teacher, newUser); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return teacher, nil
}

func (u *adminUsecase) CreateStudent(ctx context.Context, input *models.CreateStudentInput) (*models.Student, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"input": utils.Dump(input),
	})

	if err := input.Validate(); err != nil {
		return nil, ErrValidation
	}

	user := auth.GetUserFromCtx(ctx)
	if user.Role != models.RoleAdmin {
		logger.Info("unauthorized access on create student by: ", user)
		return nil, ErrUnauthorized
	}

	studentID := utils.GenerateID()
	student := &models.Student{
		ID:   studentID,
		Name: input.Name,
	}

	newUser := &models.User{
		ID:        studentID,
		Password:  input.Password,
		CreatedAt: time.Now(),
		Role:      models.RoleStudent,
	}

	if err := newUser.Encrypt(); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	if err := u.adminRepo.CreateStudent(ctx, student, newUser); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return student, nil
}

func (u *adminUsecase) CreateSubject(ctx context.Context, input *models.CreateSubjectInput) (*models.Subject, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"input": utils.Dump(input),
	})

	user := auth.GetUserFromCtx(ctx)
	if user.Role != models.RoleAdmin {
		logger.Info("unauthorized access on create subject by: ", user)
		return nil, ErrUnauthorized
	}

	_, err := u.teacherRepo.FindByID(ctx, input.TeacherID)
	switch err {
	default:
		logger.Error(err)
		return nil, ErrInternal
	case repository.ErrNotFound:
		return nil, ErrNotFound
	case nil:
		break
	}

	subject := &models.Subject{
		ID:        utils.GenerateID(),
		Name:      input.Name,
		TeacherID: input.TeacherID,
	}

	if err := u.subjectRepo.Create(ctx, subject); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return subject, nil
}
