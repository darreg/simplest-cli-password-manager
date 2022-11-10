package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type LoginDTO struct {
	Login    string
	Password string
}

func Login(
	ctx context.Context,
	client port.GRPCClientLoginSupporter,
	cliScript port.CLILoginSupporter,
) (string, error) {
	var (
		err        error
		sessionKey string
	)

	loginDTO := &LoginDTO{}
	err = cliScript.Login(ctx, loginDTO)
	if err != nil {
		return "", err
	}

	sessionKey, err = client.Login(ctx, loginDTO.Login, loginDTO.Password)
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}
