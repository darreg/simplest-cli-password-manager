package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/pkg/proto"
)

// Login calls the same GRPC method.
func (c *Client) Login(ctx context.Context, login, password string) (string, error) {
	if c.grpcClient == nil {
		return "", ErrGRPCClient
	}

	response, err := c.grpcClient.Login(ctx, &proto.LoginRequest{
		Login:    login,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return response.EncryptedSessionKey, nil
}
