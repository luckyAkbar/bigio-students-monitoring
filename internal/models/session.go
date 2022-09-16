package models

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Session struct {
	ID          int64     `json:"-"`
	AccessToken string    `json:"access_token"`
	UserID      int64     `json:"user_id"`
	Role        Role      `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiredAt   time.Time `json:"expired_at"`
}

func (s *Session) IsAccessTokenExpired() bool {
	if s == nil {
		logrus.Warn("function: IsAccessTokenExpired is called on nil Session")
		return true
	}

	now := time.Now()
	return now.After(s.ExpiredAt)
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	FindByAccessToken(ctx context.Context, token string) (*Session, error)
}
