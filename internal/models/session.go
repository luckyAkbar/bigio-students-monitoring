package models

import (
	"context"
	"time"
)

type Session struct {
	ID int64 `json:"-"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID int64 `json:"user_id"`
	Role Role `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	FindByID(ctx context.Context, sessionID int64) (*Session, error)
}