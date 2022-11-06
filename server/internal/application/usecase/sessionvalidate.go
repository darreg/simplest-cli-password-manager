package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/google/uuid"
)

func SessionValidate(
	ctx context.Context,
	encryptedSessionKey string,
	sessionLifeTime string,
	decryptor port.Decryptor,
	sessionGetter port.SessionGetter,
) (*entity.Session, error) {
	if encryptedSessionKey == "" {
		return nil, ErrInvalidSessionKey
	}

	sessionKey, err := decryptor.Decrypt(encryptedSessionKey)
	if err != nil || len(sessionKey) == 0 {
		return nil, ErrInvalidSessionKey
	}

	sessionID, err := uuid.Parse(string(sessionKey))
	if err != nil {
		return nil, ErrInternalServerError
	}

	session, err := sessionGetter.Get(ctx, sessionID)
	if err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return nil, ErrNotAuthenticated
		}
		return nil, ErrInternalServerError
	}

	duration, err := time.ParseDuration(sessionLifeTime)
	if err != nil {
		return nil, ErrInternalServerError
	}

	if session.LastSeenTime.Before(time.Now().Add(-duration)) {
		return nil, ErrNotAuthenticated
	}

	return session, nil
}
