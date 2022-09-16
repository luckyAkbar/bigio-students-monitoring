package rest

import (
	"net/http"
	"strconv"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Service) getGradeBySubjectID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		param := ctx.Param("subjectID")
		if param == "" {
			return ErrBadRequest
		}

		subjectID, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return ErrValidation
		}

		grade, err := s.studentUsecase.GetGradeBySubjectID(ctx.Request().Context(), subjectID)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":       utils.DumpIncomingContext(ctx.Request().Context()),
				"subjectID": subjectID,
			}).Error(err)
			return ErrInternal
		case usecase.ErrUnauthorized:
			return ErrUnauthorized
		case usecase.ErrNotFound:
			return ErrNotFound
		case nil:
			return ctx.JSON(http.StatusOK, grade)
		}
	}
}

func (s *Service) getAllGrades() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		grades, err := s.studentUsecase.GetAllGrade(ctx.Request().Context())
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"ctx": utils.DumpIncomingContext(ctx.Request().Context()),
			}).Error(err)
			return ErrInternal
		}

		return ctx.JSON(http.StatusOK, grades)
	}
}
