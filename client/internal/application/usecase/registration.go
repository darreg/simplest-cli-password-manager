package usecase

import (
	"context"
)

type RegistrationData struct {
	Login          string
	Password       string
	Repeatpassword string
}

func Registration(
	ctx context.Context,
	cli func(ctx context.Context, data any) error,
) (*RegistrationData, error) {
	registrationData := &RegistrationData{}

	err := cli(ctx, registrationData)
	if err != nil {
		return nil, err
	}

	if registrationData.Password != registrationData.Repeatpassword {
		return nil, ErrInvalidArgument
	}

	return registrationData, nil
}
