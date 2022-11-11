package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GetUser calls the same GRPC method.
func (c *Client) GetUser(ctx context.Context) (*model.User, error) {
	if c.grpcClient == nil {
		return nil, ErrGRPCClient
	}
	if c.sessionKey == "" {
		return nil, ErrSessionKey
	}

	md := metadata.New(map[string]string{string(port.SessionContextKey): c.sessionKey})
	ctx = metadata.NewOutgoingContext(ctx, md)
	var header metadata.MD

	response, err := c.grpcClient.GetUser(ctx, &proto.GetUserRequest{}, grpc.Header(&header))
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:    response.UserId,
		Name:  response.Name,
		Login: response.Login,
	}, nil
}
