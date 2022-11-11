package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SetEntry calls the same GRPC method.
func (c *Client) SetEntry(ctx context.Context, typeID, name, mdata string, data []byte) error {
	if c.grpcClient == nil {
		return ErrGRPCClient
	}
	if c.sessionKey == "" {
		return ErrSessionKey
	}

	md := metadata.New(map[string]string{string(port.SessionContextKey): c.sessionKey})
	ctx = metadata.NewOutgoingContext(ctx, md)
	var header metadata.MD

	_, err := c.grpcClient.SetEntry(ctx, &proto.SetEntryRequest{
		TypeId:   typeID,
		Name:     name,
		Metadata: mdata,
		Data:     data,
	}, grpc.Header(&header))
	if err != nil {
		return err
	}

	return nil
}
