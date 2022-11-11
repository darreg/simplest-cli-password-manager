package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/pkg/proto"
)

// Registration calls the same GRPC method.
func (c *Client) Registration(ctx context.Context, name, login, password string) (string, error) {
	if c.grpcClient == nil {
		return "", ErrGRPCClient
	}

	response, err := c.grpcClient.Registration(ctx, &proto.RegistrationRequest{
		Name:     name,
		Login:    login,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return response.EncryptedSessionKey, nil
}
