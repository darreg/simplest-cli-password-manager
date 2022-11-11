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

// GetEntry gets an entry.
func (c *Collection) GetEntry(ctx context.Context, in *proto.GetEntryRequest) (*proto.GetEntryResponse, error) {
	entry, err := usecase.GetEntry(ctx, in.EntryId, c.a.Encryptor, c.a.EntryRepository, c.a.UserRepository)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrEntryNotFound):
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	return &proto.GetEntryResponse{
		EntryId:   entry.ID.String(),
		TypeId:    entry.TypeID.String(),
		Name:      entry.Name,
		Metadata:  entry.Metadata,
		Data:      entry.Data,
		CreatedAt: timestamppb.New(*entry.CreatedAt),
		UpdatedAt: timestamppb.New(*entry.UpdatedAt),
	}, nil
}
