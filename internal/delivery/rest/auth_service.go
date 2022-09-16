package rest

import (
	"net/http"

	"github.com/kumparan/go-utils"
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/usecase"
	"github.com/sirupsen/logrus"
)

func (s *Service) LoginByIDAndPassword() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := &models.LoginByIDAndPasswordInput{}
		if err := ctx.Bind(input); err != nil {
			return ErrBadRequest
		}

		session, err := s.authUsecase.LoginByIDAndPassword(ctx.Request().Context(), input.ID, input.Password)
		switch err {
		default:
			logrus.WithFields(logrus.Fields{
				"ctx":   utils.DumpIncomingContext(ctx.Request().Context()),
				"input": utils.Dump(input),
			}).Error(err)
			return ErrInternal
		case usecase.ErrNotFound:
			return ErrNotFound
		case usecase.ErrUnauthorized:
			return ErrUnauthorized
		case nil:
			return ctx.JSON(http.StatusOK, session)
		}
	}
}
