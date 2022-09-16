package usecase

import "errors"

var (
	ErrInternal = errors.New("internal error")
	ErrNotFound = errors.New("not found")
	ErrUnauthorized = errors.New("authorization failed")
)