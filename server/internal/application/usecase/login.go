package usecase

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type Credential struct {
	Login    string
	Password string
}

func Login(
	ctx context.Context,
	cred Credential,
	userRepository port.UserByCredentialGetter,
	hasher port.PasswordHasher,
) (*entity.User, error) {
	user, err := userRepository.GetByCredential(ctx, cred.Login, hasher.Hash(cred.Password))
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrNotAuthenticated
		}
		return nil, ErrInternalServerError
	}

	return user, nil
}
