package usecase

import "errors"

var (
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrInternalError     = errors.New("internal error")
	ErrIncorrectPassword = errors.New("passwords don't match")
)
