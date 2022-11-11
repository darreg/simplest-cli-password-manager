package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GetEntry calls the same GRPC method.
func (c *Client) GetEntry(ctx context.Context, entryID string) (*model.Entry, error) {
	if c.grpcClient == nil {
		return nil, ErrGRPCClient
	}
	if c.sessionKey == "" {
		return nil, ErrSessionKey
	}

	md := metadata.New(map[string]string{string(port.SessionContextKey): c.sessionKey})
	ctx = metadata.NewOutgoingContext(ctx, md)
	var header metadata.MD

	response, err := c.grpcClient.GetEntry(ctx, &proto.GetEntryRequest{
		EntryId: entryID,
	}, grpc.Header(&header))
	if err != nil {
		return nil, err
	}

	return &model.Entry{
		ID:       response.EntryId,
		TypeID:   response.TypeId,
		Name:     response.Name,
		Metadata: response.Metadata,
		Data:     response.Data,
	}, nil
}
