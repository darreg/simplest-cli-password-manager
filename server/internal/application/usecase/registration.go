package usecase

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/google/uuid"
)

type RegistrationData struct {
	Login    string
	Password string
}

func Registration(
	ctx context.Context,
	regData RegistrationData,
	userRepository port.UserRegistrator,
	hasher port.PasswordHasher,
) (*entity.User, error) {
	user, err := userRepository.GetByLogin(ctx, regData.Login)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return nil, err
	}

	if user != nil {
		return nil, ErrLoginAlreadyUse
	}

	user = &entity.User{
		ID:           uuid.New(),
		Login:        regData.Login,
		PasswordHash: hasher.Hash(regData.Password),
	}
	err = userRepository.Add(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
