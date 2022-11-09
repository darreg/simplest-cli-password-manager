package interceptor

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Auth interceptor for authenticates the user.
func Auth(sessionLifeTime string, sessionRepository port.SessionRepository, decryptor port.Decryptor) func(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	unaryHandler grpc.UnaryHandler,
) (interface{}, error) {
	fn := func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		unaryHandler grpc.UnaryHandler,
	) (interface{}, error) {
		if info.FullMethod == "/proto.App/Registration" ||
			info.FullMethod == "/proto.App/Login" ||
			info.FullMethod == "/proto.App/GetAllTypes" {
			return unaryHandler(ctx, req)
		}

		var encryptedSessionKey string

		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			values := md.Get(string(port.SessionContextKey))
			if len(values) > 0 {
				encryptedSessionKey = values[0]
			}
		}

		session, err := usecase.SessionValidate(ctx, encryptedSessionKey, sessionLifeTime, decryptor, sessionRepository)
		if err != nil {
			switch {
			case errors.Is(err, usecase.ErrNotAuthenticated):
				return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
			case errors.Is(err, usecase.ErrInvalidSessionKey):
				return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
			default:
				return nil, status.Error(codes.Internal, codes.Internal.String())
			}
		}

		ctx = context.WithValue(ctx, port.SessionContextKey, session)

		return unaryHandler(ctx, req)
	}

	return fn
}
