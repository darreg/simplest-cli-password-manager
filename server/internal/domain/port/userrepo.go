package port

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/google/uuid"
)

type UserGetter interface {
	Get(ctx context.Context, userID uuid.UUID) (*entity.User, error)
}

type UserByLoginGetter interface {
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
}

type UserByCredentialGetter interface {
	GetByCredential(ctx context.Context, login, passwordHash string) (*entity.User, error)
}

type UserAdder interface {
	Add(ctx context.Context, user *entity.User) error
}

type UserChanger interface {
	Change(ctx context.Context, user *entity.User) error
}

type UserPasswordChanger interface {
	ChangePassword(ctx context.Context, user *entity.User) error
}

type UserRemover interface {
	Remove(ctx context.Context, userID uuid.UUID) error
}

type UserRegistrator interface {
	UserByLoginGetter
	UserAdder
}

type UserRepository interface {
	UserGetter
	UserByLoginGetter
	UserByCredentialGetter
	UserAdder
	UserChanger
	UserPasswordChanger
	UserRemover
}
