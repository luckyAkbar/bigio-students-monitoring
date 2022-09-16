package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
)

type Service struct {
	group        *echo.Group
	authUsecase  models.AuthUsecase
	adminUsecase models.AdminUsecase
}

func NewRESTService(group *echo.Group, authUsecase models.AuthUsecase, adminUsecase models.AdminUsecase) {
	service := &Service{
		group:        group,
		authUsecase:  authUsecase,
		adminUsecase: adminUsecase,
	}

	service.InitRoutes()
}

func (s *Service) InitRoutes() {
	s.initAuthService()
	s.initAdminService()
}

func (s *Service) initAuthService() {
	s.group.POST("/auth/login/", s.loginByIDAndPassword())
}

func (s *Service) initAdminService() {
	s.group.POST("/admin/create/teacher/", s.createTeacher())
	s.group.POST("/admin/create/student/", s.createStudent())
	s.group.POST("/admin/create/subject/", s.createSubject())
}
