package client

import (
	"context"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/cli"
	"github.com/alrund/yp-2-project/client/pkg/proto"
)

func (c *Client) Registration(ctx context.Context) error {
	if c.GRPCClient == nil {
		return ErrGRPCClient
	}

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

	c.EncryptedSessionKey = response.EncryptedSessionKey

	return nil
}
