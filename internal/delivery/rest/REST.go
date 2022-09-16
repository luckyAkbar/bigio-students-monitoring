package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
)

type Service struct {
	group       *echo.Group
	authUsecase models.AuthUsecase
}

func NewRESTService(group *echo.Group, authUsecase models.AuthUsecase) {
	service := &Service{
		group:       group,
		authUsecase: authUsecase,
	}

	service.InitRoutes()
}

func (s *Service) InitRoutes() {
	s.initAuthService()
}

func (s *Service) initAuthService() {
	s.group.POST("/auth/login/", s.LoginByIDAndPassword())
}
