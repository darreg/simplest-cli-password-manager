package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type LoginDTO struct {
	Login    string
	Password string
}

// Login processes the authorization form.
func Login(
	ctx context.Context,
	client port.GRPCClientLoginSupporter,
	cliScript port.CLILoginSupporter,
) (string, error) {
	var (
		err        error
		sessionKey string
	)

	dto, err := cliScript.Login(ctx)
	if err != nil {
		return "", err
	}

	loginDTO, ok := dto.(*LoginDTO)
	if !ok {
		return "", ErrInternalError
	}

	sessionKey, err = client.Login(ctx, loginDTO.Login, loginDTO.Password)
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}
