package handler

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAllTypes gets all types.
func (c *Collection) GetAllTypes(
	ctx context.Context,
	in *proto.GetAllTypesRequest,
) (*proto.GetAllTypesResponse, error) {
	types, err := usecase.GetAllTypes(ctx, c.a.TypeRepository)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrTypeNotFound):
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	responseTypes := make([]*proto.GetAllTypesResponse_Type, len(types))
	for i, tp := range types {
		responseTypes[i] = &proto.GetAllTypesResponse_Type{
			TypeId:   tp.ID.String(),
			Name:     tp.Name,
			IsBinary: tp.IsBinary,
		}
	}

	return &proto.GetAllTypesResponse{Types: responseTypes}, nil
}
