package usecase

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/google/uuid"
)

// GetEntry gets entry by ID.
func GetEntry(
	ctx context.Context,
	requestedEntryID string,
	decryptor port.Decryptor,
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
		if errors.Is(err, ErrEntryNotFound) {
			return nil, ErrEntryNotFound
		}
		return nil, ErrInternalServerError
	}

	entry.Data, err = decryptor.Decrypt(string(entry.Data))
	if err != nil {
		return nil, ErrInternalServerError
	}

	return entry, nil
}
