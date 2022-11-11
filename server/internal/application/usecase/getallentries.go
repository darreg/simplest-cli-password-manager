package usecase

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

// GetAllEntries gets all user entries.
func GetAllEntries(
	ctx context.Context,
	entryRepository port.EntryAllByUserGetter,
	userRepository port.UserGetter,
) ([]*entity.Entry, error) {
	contextSession := ctx.Value(port.SessionContextKey)
	session, ok := contextSession.(*entity.Session)
	if !ok {
		return nil, ErrIncorrectSession
	}

	user, err := userRepository.Get(ctx, session.UserID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	entries, err := entryRepository.GetAllByUser(ctx, user)
	if err != nil {
		if errors.Is(err, ErrEntryNotFound) {
			return nil, ErrEntryNotFound
		}
		return nil, ErrInternalServerError
	}

	return entries, nil
}
