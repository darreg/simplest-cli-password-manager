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

func Login(
	ctx context.Context,
	cred Credential,
	hasher port.PasswordHasher,
	userRepository port.UserByCredentialGetter,
	sessionRepository port.SessionAdder,
) (*entity.Session, error) {
	user, err := userRepository.GetByCredential(ctx, cred.Login, hasher.Hash(cred.Password))
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrNotAuthenticated
		}
		return nil, ErrInternalServerError
	}

	loginTime := time.Now()
	session := entity.NewSession(user.ID, &loginTime, &loginTime)

	err = sessionRepository.Add(ctx, session)
	if err != nil {
		return nil, ErrInternalServerError
	}

	return session, nil
}
