package client

import (
	"errors"

	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/pkg/proto"
)

var (
	ErrGRPCClient = errors.New("no GRPS client")
	ErrSessionKey = errors.New("no session")
)

// Client wrapper over GRPC client.
type Client struct {
	grpcClient proto.AppClient
	sessionKey string
	types      []*model.Type
}

func New() *Client {
	return &Client{}
}

func (c *Client) SetGRPCClient(client any) error {
	grpcClient, ok := client.(proto.AppClient)
	if !ok {
		return usecase.ErrInvalidArgument
	}
	c.grpcClient = grpcClient
	return nil
}

func (c *Client) SetSessionKey(sessionKey string) error {
	if sessionKey == "" {
		return usecase.ErrInvalidArgument
	}
	c.sessionKey = sessionKey
	return nil
}

func (c *Client) IsEmptySessionKey() bool {
	return c.sessionKey == ""
}
