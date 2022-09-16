package rest

import (
	"net/http"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Service) gradeByStudentID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &models.CreateGradeInput{}
		if err := ctx.Bind(input); err != nil {
			return ErrBadRequest
		}

		logrus.Info(input)

		mark, err := models.ConvertStringToMark(input.Mark)
		if err != nil {
			return ErrValidation
		}

		input.Mark = mark

		grade, err := s.teacherUsecase.GradeByStudentID(ctx.Request().Context(), input)
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
		case usecase.ErrNotFound:
			return ErrNotFound
		case nil:
			return ctx.JSON(http.StatusOK, grade)
		}
	}
}
