package client

import (
	"context"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/cli"
	"github.com/alrund/yp-2-project/client/pkg/proto"
)

const LoginIndex int = 0
const RegistrationIndex int = 1

func (c *Client) Login(ctx context.Context) error {
	if c.GRPCClient == nil {
		return ErrGRPCClient
	}

	var (
		encryptedSessionKey string
		loginMethods        = []string{LoginIndex: "Login", RegistrationIndex: "Registration"}
		loginMethodIndex    int
	)
	err := cli.SelectLoginMethod(ctx, loginMethods, &loginMethodIndex)
	if err != nil {
		return err
	}

	switch loginMethodIndex {
	case LoginIndex:
		credential, err := usecase.Login(ctx, cli.Login)
		if err != nil {
			return err
		}

		response, err := c.GRPCClient.Login(ctx, &proto.LoginRequest{
			Login:    credential.Login,
			Password: credential.Password,
		})
		if err != nil {
			return err
		}

		encryptedSessionKey = response.EncryptedSessionKey

	case RegistrationIndex:
		registrationData, err := usecase.Registration(ctx, cli.Registration)
		if err != nil {
			return err
		}

		response, err := c.GRPCClient.Registration(ctx, &proto.RegistrationRequest{
			Login:    registrationData.Login,
			Password: registrationData.Password,
		})
		if err != nil {
			return err
		}

		encryptedSessionKey = response.EncryptedSessionKey
	}

	c.EncryptedSessionKey = encryptedSessionKey

	return nil
}
