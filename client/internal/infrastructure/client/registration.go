package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/pkg/proto"
)

func (c *Client) Registration(ctx context.Context, login, password string) (string, error) {
	if c.grpcClient == nil {
		return "", ErrGRPCClient
	}

	response, err := c.grpcClient.Registration(ctx, &proto.RegistrationRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return response.EncryptedSessionKey, nil
}
