package rest

import (
	"net/http"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Service) createTeacher() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &models.CreateTeacherInput{}
		if err := ctx.Bind(input); err != nil {
			return ErrBadRequest
		}

		teacher, err := s.adminUsecase.CreateTeacher(ctx.Request().Context(), input)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(ctx.Request().Context()),
				"input": utils.Dump(input),
			}).Error(err)

			return ErrInternal
		case usecase.ErrUnauthorized:
			return ErrUnauthorized
		case usecase.ErrValidation:
			return ErrValidation
		case nil:
			return ctx.JSON(http.StatusOK, teacher)
		}
	}
}

func (s *Service) createStudent() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &models.CreateStudentInput{}
		if err := ctx.Bind(input); err != nil {
			return ErrBadRequest
		}

		student, err := s.adminUsecase.CreateStudent(ctx.Request().Context(), input)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(ctx.Request().Context()),
				"input": utils.Dump(input),
			}).Error(err)

			return ErrInternal
		case usecase.ErrUnauthorized:
			return ErrUnauthorized
		case usecase.ErrValidation:
			return ErrValidation
		case nil:
			return ctx.JSON(http.StatusOK, student)
		}
	}
}

func (s *Service) createSubject() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &models.CreateSubjectInput{}
		if err := ctx.Bind(input); err != nil {
			return ErrBadRequest
		}

		subject, err := s.adminUsecase.CreateSubject(ctx.Request().Context(), input)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(ctx.Request().Context()),
				"input": utils.Dump(input),
			}).Error(err)

			return ErrInternal

		case usecase.ErrUnauthorized:
			return ErrUnauthorized
		case usecase.ErrNotFound:
			return ErrNotFound
		case nil:
			return ctx.JSON(http.StatusOK, subject)
		}
	}
}
