package handler

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser gets a user.
func (c *Collection) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := usecase.GetUser(ctx, c.a.UserRepository)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	return &proto.GetUserResponse{
		UserId: user.ID.String(),
		Name:   user.Name,
		Login:  user.Login,
	}, nil
}
