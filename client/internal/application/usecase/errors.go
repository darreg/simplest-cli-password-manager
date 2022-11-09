package usecase

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInternalError   = errors.New("internal error")
)
