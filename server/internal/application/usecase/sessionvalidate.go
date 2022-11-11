package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/google/uuid"
)

// SessionValidate validates of the session key.
func SessionValidate(
	ctx context.Context,
	encryptedSessionKey string,
	sessionLifeTime string,
	decryptor port.Decryptor,
	sessionRepository port.SessionRefresher,
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

	session, err := sessionRepository.Get(ctx, sessionID)
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

	nowTime := time.Now()
	session.LastSeenTime = &nowTime
	err = sessionRepository.Change(ctx, session)
	if err != nil {
		return nil, ErrInternalServerError
	}

	return session, nil
}
