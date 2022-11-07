package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type RegistrationData struct {
	Login    string
	Password string
}

func Registration(
	ctx context.Context,
	regData RegistrationData,
	hasher port.PasswordHasher,
	userRepository port.UserRegistrator,
	sessionRepository port.SessionAdder,
) (*entity.Session, error) {
	user, err := userRepository.GetByLogin(ctx, regData.Login)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return nil, ErrInternalServerError
	}

	if user != nil {
		return nil, ErrLoginAlreadyUse
	}

	user = entity.NewUser(regData.Login, hasher.Hash(regData.Password))
	err = userRepository.Add(ctx, user)
	if err != nil {
		return nil, ErrInternalServerError
	}

	nowTime := time.Now()
	session := entity.NewSession(user.ID, &nowTime, &nowTime)
	err = sessionRepository.Add(ctx, session)
	if err != nil {
		return nil, ErrInternalServerError
	}

	return session, nil
}
