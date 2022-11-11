package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type RegistrationDTO struct {
	Name           string
	Login          string
	Password       string
	RepeatPassword string
}

// Registration processes the registration form.
func Registration(
	ctx context.Context,
	client port.GRPCClientRegistrationSupporter,
	cliScript port.CLIRegistrationSupporter,
) (string, error) {
	var (
		err        error
		sessionKey string
	)

	dto, err := cliScript.Registration(ctx)
	if err != nil {
		return "", err
	}

	registrationDTO, ok := dto.(*RegistrationDTO)
	if !ok {
		return "", ErrInternalError
	}

	if registrationDTO.Password != registrationDTO.RepeatPassword {
		return "", ErrIncorrectPassword
	}

	sessionKey, err = client.Registration(
		ctx,
		registrationDTO.Name,
		registrationDTO.Login,
		registrationDTO.Password,
	)
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}
