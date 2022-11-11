package handler

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetAllEntries gets all entries.
func (c *Collection) GetAllEntries(
	ctx context.Context,
	in *proto.GetAllEntriesRequest,
) (*proto.GetAllEntriesResponse, error) {
	entries, err := usecase.GetAllEntries(ctx, c.a.EntryRepository, c.a.UserRepository)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrEntryNotFound):
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	responseEntries := make([]*proto.GetAllEntriesResponse_Entry, len(entries))
	for i, entry := range entries {
		responseEntries[i] = &proto.GetAllEntriesResponse_Entry{
			EntryId:   entry.ID.String(),
			TypeId:    entry.TypeID.String(),
			Name:      entry.Name,
			CreatedAt: timestamppb.New(*entry.CreatedAt),
			UpdatedAt: timestamppb.New(*entry.UpdatedAt),
		}
	}

	return &proto.GetAllEntriesResponse{Entries: responseEntries}, nil
}
