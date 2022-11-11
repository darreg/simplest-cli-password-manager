package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

// GetUser gets user.
func GetUser(
	ctx context.Context,
	userRepository port.UserGetter,
) (*entity.User, error) {
	contextSession := ctx.Value(port.SessionContextKey)
	session, ok := contextSession.(*entity.Session)
	if !ok {
		return nil, ErrIncorrectSession
	}

	user, err := userRepository.Get(ctx, session.UserID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
