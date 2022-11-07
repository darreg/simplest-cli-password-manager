package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/google/uuid"
)

func GetEntry(
	ctx context.Context,
	requestedEntryID string,
	entryRepository port.EntryOneWithUserGetter,
	userRepository port.UserGetter,
) (*entity.Entry, error) {
	contextSession := ctx.Value(port.SessionContextKey)
	session, ok := contextSession.(*entity.Session)
	if !ok {
		return nil, ErrIncorrectSession
	}

	user, err := userRepository.Get(ctx, session.UserID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	entryID, err := uuid.Parse(requestedEntryID)
	if err != nil {
		return nil, ErrInvalidArgument
	}

	entry, err := entryRepository.GetOneWithUser(ctx, entryID, user)
	if err != nil {
		return nil, err
	}

	return entry, nil
}
