package client

import (
	"errors"

	"github.com/alrund/yp-2-project/client/internal/application/app"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/pkg/proto"
)

var ErrGRPCClient = errors.New("no GRPS client set")

type Client struct {
	a                   *app.App
	GRPCClient          proto.AppClient
	EncryptedSessionKey string
}

func New(a *app.App) *Client {
	return &Client{a: a}
}

func (c *Client) SetGRPCClient(client any) error {
	GRPCClient, ok := client.(proto.AppClient)
	if !ok {
		return usecase.ErrInvalidArgument
	}
	c.GRPCClient = GRPCClient
	return nil
}
