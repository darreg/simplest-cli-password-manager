package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type Credential struct {
	Login    string
	Password string
}

// Login authorizes the user and returns an encrypted session key.
func Login(
	ctx context.Context,
	cred Credential,
	hasher port.PasswordHasher,
	encryptor port.Encryptor,
	userRepository port.UserByCredentialGetter,
	sessionRepository port.SessionAdder,
) (string, error) {
	user, err := userRepository.GetByCredential(ctx, cred.Login, hasher.Hash(cred.Password))
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return "", ErrNotAuthenticated
		}
		return "", ErrInternalServerError
	}

	loginTime := time.Now()
	session := entity.NewSession(user.ID, &loginTime, &loginTime)

	err = sessionRepository.Add(ctx, session)
	if err != nil {
		return "", ErrInternalServerError
	}

	encryptedSessionKey, err := encryptor.Encrypt([]byte(session.ID.String()))
	if err != nil {
		return "", ErrInternalServerError
	}

	return encryptedSessionKey, nil
}
