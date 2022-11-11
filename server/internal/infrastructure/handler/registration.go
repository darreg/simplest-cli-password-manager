package handler

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Registration registers a new user.
func (c *Collection) Registration(
	ctx context.Context,
	in *proto.RegistrationRequest,
) (*proto.RegistrationResponse, error) {
	encryptedSessionKey, err := usecase.Registration(
		ctx,
		usecase.RegistrationData{
			Name:     in.Name,
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
		case errors.Is(err, usecase.ErrLoginAlreadyUse):
			return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
		default:
			return nil, status.Error(codes.Internal, codes.Internal.String())
		}
	}

	return &proto.RegistrationResponse{EncryptedSessionKey: encryptedSessionKey}, nil
}
