package console

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luckyAkbar/bigio-students-monitoring/auth"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/config"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/db"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/delivery/rest"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/repository"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start the serveer",
	Long:  "start the server",
	Run:   server,
}

func init() {
	RootCMD.AddCommand(serverCmd)
}

func server(cmd *cobra.Command, args []string) {
	setupLogger()
	db.InitializePostgresConn()

	sqlDB, err := db.PostgresDB.DB()
	if err != nil {
		logrus.Fatal("unable to start server. reason: ", err.Error())
	}

	defer sqlDB.Close()

	sessionRepo := repository.NewSessionRepo(db.PostgresDB)
	userRepo := repository.NewUserRepository(db.PostgresDB)
	adminRepo := repository.NewAdminRepository(db.PostgresDB)
	teacherRepo := repository.NewTeacherRepository(db.PostgresDB)
	subjectRepo := repository.NewSubjectRepository(db.PostgresDB)
	studentRepo := repository.NewStudentRepository(db.PostgresDB)
	gradeRepo := repository.NewGradeRepository(db.PostgresDB)

	authUsecase := usecase.NewAuthUsecase(sessionRepo, userRepo)
	adminUsecase := usecase.NewAdminUsecase(adminRepo, teacherRepo, subjectRepo)
	teacherUsecase := usecase.NewTeacherUsecase(studentRepo, subjectRepo, gradeRepo)

	authMiddleware := auth.NewMiddleware(sessionRepo, userRepo)

	HTTPServer := echo.New()

	HTTPServer.Pre(middleware.AddTrailingSlash())

	HTTPServer.Use(middleware.Logger())
	HTTPServer.Use(authMiddleware.UserSessionMiddleware())
	HTTPServer.Use(authMiddleware.RejectUnauthorizedRequest())

	RESTGroup := HTTPServer.Group("rest")

	rest.NewRESTService(RESTGroup, authUsecase, adminUsecase, teacherUsecase)

	if err := HTTPServer.Start(fmt.Sprintf(":%s", config.ServerPort())); err != nil {
		logrus.Fatal("unable to start server. reason: ", err.Error())
	}

	logrus.Info("Server running on port: ", config.ServerPort())
}
