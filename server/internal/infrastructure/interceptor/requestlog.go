package interceptor

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"google.golang.org/grpc"
)

// RequestLog logs all requests.
func RequestLog(logger port.Logger) func(
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
		var userID string
		contextSession := ctx.Value(port.SessionContextKey)
		session, ok := contextSession.(*entity.Session)
		if ok {
			userID = session.UserID.String()
		}

		logger.Info("request",
			"method", info.FullMethod,
			"userID", userID,
		)

		return unaryHandler(ctx, req)
	}

	return fn
}
