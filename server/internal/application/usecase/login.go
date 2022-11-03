package usecase

import (
	"context"

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
		return nil, err
	}

	return user, nil
}
