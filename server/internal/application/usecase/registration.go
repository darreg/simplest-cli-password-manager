package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type RegistrationData struct {
	Name     string
	Login    string
	Password string
}

// Registration adds the user and returns an encrypted session key.
func Registration(
	ctx context.Context,
	regData RegistrationData,
	hasher port.PasswordHasher,
	encryptor port.Encryptor,
	userRepository port.UserRegistrator,
	sessionRepository port.SessionAdder,
) (string, error) {
	user, err := userRepository.GetByLogin(ctx, regData.Login)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return "", ErrInternalServerError
	}

	if user != nil {
		return "", ErrLoginAlreadyUse
	}

	user = entity.NewUser(regData.Name, regData.Login, hasher.Hash(regData.Password))
	err = userRepository.Add(ctx, user)
	if err != nil {
		return "", ErrInternalServerError
	}

	nowTime := time.Now()
	session := entity.NewSession(user.ID, &nowTime, &nowTime)
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
