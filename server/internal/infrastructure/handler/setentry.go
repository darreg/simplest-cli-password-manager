package handler

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SetEntry adds an entry.
func (c *Collection) SetEntry(ctx context.Context, in *proto.SetEntryRequest) (*proto.SetEntryResponse, error) {
	err := usecase.SetEntry(
		ctx,
		&usecase.SetEntryDTO{
			TypeID:   in.TypeId,
			Name:     in.Name,
			Metadata: in.Metadata,
			Data:     in.Data,
		},
		c.a.Encryptor,
		c.a.EntryRepository,
		c.a.UserRepository,
		c.a.TypeRepository,
	)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidArgument):
			return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	return &proto.SetEntryResponse{}, nil
}
