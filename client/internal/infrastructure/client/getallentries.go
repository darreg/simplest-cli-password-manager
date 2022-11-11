package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GetAllEntries calls the same GRPC method.
func (c *Client) GetAllEntries(ctx context.Context) ([]*model.Entry, error) {
	if c.grpcClient == nil {
		return nil, ErrGRPCClient
	}
	if c.sessionKey == "" {
		return nil, ErrSessionKey
	}

	md := metadata.New(map[string]string{string(port.SessionContextKey): c.sessionKey})
	ctx = metadata.NewOutgoingContext(ctx, md)
	var header metadata.MD

	response, err := c.grpcClient.GetAllEntries(ctx, &proto.GetAllEntriesRequest{}, grpc.Header(&header))
	if err != nil {
		return nil, err
	}

	entries := make([]*model.Entry, len(response.Entries))
	for i, entry := range response.Entries {
		entries[i] = &model.Entry{
			ID:     entry.EntryId,
			TypeID: entry.TypeId,
			Name:   entry.Name,
		}
	}

	return entries, nil
}
