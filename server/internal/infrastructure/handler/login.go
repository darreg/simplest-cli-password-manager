package handler

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login authenticates the user with username and password.
func (c *Collection) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	encryptedSessionKey, err := usecase.Login(
		ctx,
		usecase.Credential{
			Login:    in.Login,
			Password: in.Password,
		},
		c.a.Hasher,
		c.a.Encryptor,
		c.a.UserRepository,
		c.a.SessionRepository,
	)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrNotAuthenticated):
			return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	return &proto.LoginResponse{EncryptedSessionKey: encryptedSessionKey}, nil
}
