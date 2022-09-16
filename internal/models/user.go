package models

import (
	"context"
	"time"

	"github.com/luckyAkbar/bigio-students-monitoring/internal/helper"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID int64 `json:"id"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Role Role `json:"role"`
}

func (u *User) Encrypt() error {
	password, err := helper.HashString(u.Password)
	if err != nil {
		logrus.Error(err)
		return err
	}

	u.Password = password
	return nil;
}

type CreateUserInput struct {
	Password string `json:"password"`
	Role Role `json:"role"`	
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
}