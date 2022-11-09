package usecase

import (
	"context"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type LoginDTO struct {
	Login    string
	Password string
}

type RegistrationDTO struct {
	Login          string
	Password       string
	RepeatPassword string
}

const (
	LoginIndex = iota
	RegistrationIndex
)

func Login(
	ctx context.Context,
	client port.GRPCClientLoginMethodSupporter,
	cliScript port.CLILoginMethodSupporter,
) error {
	var (
		sessionKey       string
		loginMethods     = []string{LoginIndex: "Login", RegistrationIndex: "Registration"}
		loginMethodIndex int
	)

	err := cliScript.SelectLoginMethod(ctx, loginMethods, &loginMethodIndex)
	if err != nil {
		return err
	}

	switch loginMethodIndex {
	case LoginIndex:
		loginDTO := &LoginDTO{}
		err = cliScript.Login(ctx, loginDTO)
		if err != nil {
			return err
		}

		sessionKey, err = client.Login(ctx, loginDTO.Login, loginDTO.Password)
		if err != nil {
			return err
		}

	case RegistrationIndex:
		registrationDTO := &RegistrationDTO{}
		err = cliScript.Registration(ctx, registrationDTO)
		if err != nil {
			return err
		}

		if registrationDTO.Password != registrationDTO.RepeatPassword {
			return ErrInvalidArgument
		}

		sessionKey, err = client.Registration(ctx, registrationDTO.Login, registrationDTO.Password)
		if err != nil {
			return err
		}
	}

	err = client.SetSessionKey(sessionKey)
	if err != nil {
		return err
	}

	return nil
}
