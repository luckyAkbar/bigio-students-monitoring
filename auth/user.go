package auth

import (
	"context"

	"github.com/luckyAkbar/bigio-students-monitoring/internal/models"
)

type contextKey string

const (
	userCtxKey           contextKey = "github.com/luckyAkbar/bigio-students-monitoring/internal/auth.User"
	_headerAuthorization string     = "Authorization"
	_authScheme          string     = "Bearer"
)

// User define any data related for identifiying user
type User struct {
	ID   int64       `json:"id"`
	Role models.Role `json:"role"`
}

// SetUserToCtx self explained
func SetUserToCtx(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

// GetUserFromCtx self explained
func GetUserFromCtx(ctx context.Context) *User {
	user, ok := ctx.Value(userCtxKey).(User)
	if !ok {
		return nil
	}

	return &user
}
