package models

import (
	"context"
)

type Role = string

var (
	RoleAdmin Role = "ADMIN"
	RoleStudent Role = "STUDENT"
	RoleTeacher Role = "TEACHER"
)

type AuthUsecase interface {
	LoginByIDAndRole(ctx context.Context, id int64, role Role) (*Session, error)
}