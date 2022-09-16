package console

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/kumparan/go-utils"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/db"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var initAdminCMD = &cobra.Command{
	Use:   "init-admin",
	Short: "init admin user",
	Long:  "create a new admin-level user",
	Run:   admin,
}

func init() {
	RootCMD.AddCommand(initAdminCMD)
}

func admin(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		logrus.Fatal("Args - 1 is required as admin name and Args - 2 is required as admin password")
	}

	name := strings.TrimSpace(args[0])
	password := strings.TrimSpace(args[1])
	if password == "" || name == "" {
		logrus.Fatal("username and password is required is required as admin credentials")
	}

	now := time.Now()

	user := &models.User{
		ID:        utils.GenerateID(),
		Password:  password,
		CreatedAt: now,
		Role:      models.RoleAdmin,
	}

	if err := user.Encrypt(); err != nil {
		logrus.Fatal("failed to encrypt user password: ", err)
	}

	admin := &models.Admin{
		ID:   user.ID,
		Name: name,
	}

	db.InitializePostgresConn()

	ctx := context.Background()

	db.PostgresDB.Transaction(func(tx *gorm.DB) error {
		adminRepo := repository.NewAdminRepository(tx)
		userRepo := repository.NewUserRepository(tx)

		if err := adminRepo.Create(ctx, admin); err != nil {
			logrus.Error("failed to create admin: ", err)
			return err
		}

		if err := userRepo.Create(ctx, user); err != nil {
			logrus.Error("failed to create admin: ", err)
			return err
		}

		return nil
	})

	fmt.Printf("ID: %d\n", user.ID)
}
