package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
)

type Service struct {
	group          *echo.Group
	authUsecase    models.AuthUsecase
	adminUsecase   models.AdminUsecase
	teacherUsecase models.TeacherUsecase
	studentUsecase models.StudentUsecase
}

func NewRESTService(group *echo.Group, authUsecase models.AuthUsecase, adminUsecase models.AdminUsecase, teacherUsecase models.TeacherUsecase, studentUsecase models.StudentUsecase) {
	service := &Service{
		group:          group,
		authUsecase:    authUsecase,
		adminUsecase:   adminUsecase,
		teacherUsecase: teacherUsecase,
		studentUsecase: studentUsecase,
	}

	service.InitRoutes()
}

func (s *Service) InitRoutes() {
	s.initAuthService()
	s.initAdminService()
	s.initTeacherService()
	s.initStudentService()
}

func (s *Service) initAuthService() {
	s.group.POST("/auth/login/", s.loginByIDAndPassword())
}

func (s *Service) initAdminService() {
	s.group.POST("/admin/create/teacher/", s.createTeacher())
	s.group.POST("/admin/create/student/", s.createStudent())
	s.group.POST("/admin/create/subject/", s.createSubject())
}

func (s *Service) initTeacherService() {
	s.group.POST("/teacher/create/grade/", s.gradeByStudentID())
}

func (s *Service) initStudentService() {
	s.group.GET("/student/grade/", s.getAllGrades())
	s.group.GET("/student/grade/:subjectID/", s.getGradeBySubjectID())
}
