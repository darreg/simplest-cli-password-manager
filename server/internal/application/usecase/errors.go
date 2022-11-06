package usecase

import (
	"errors"
)

var (
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrInvalidRequestFormat = errors.New("invalid request format")
	ErrNotAuthenticated     = errors.New("not authenticated")
	ErrInternalServerError  = errors.New("internal server error")
	ErrInvalidSessionKey    = errors.New("invalid session key")
	ErrIncorrectSession     = errors.New("incorrect session")
	ErrLoginAlreadyUse      = errors.New("the login is already in use")
	ErrUserNotFound         = errors.New("user not found")
	ErrTypeNotFound         = errors.New("type not found")
	ErrEntryNotFound        = errors.New("entry not found")
	ErrSessionNotFound      = errors.New("session not found")
)
