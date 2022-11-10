package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type RegistrationDTO struct {
	Login          string
	Password       string
	RepeatPassword string
}

func Registration(
	ctx context.Context,
	client port.GRPCClientRegistrationSupporter,
	cliScript port.CLIRegistrationSupporter,
) (string, error) {
	var (
		err        error
		sessionKey string
	)

	registrationDTO := &RegistrationDTO{}
	err = cliScript.Registration(ctx, registrationDTO)
	if err != nil {
		return "", err
	}

	if registrationDTO.Password != registrationDTO.RepeatPassword {
		return "", ErrInvalidArgument
	}

	sessionKey, err = client.Registration(ctx, registrationDTO.Login, registrationDTO.Password)
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}
