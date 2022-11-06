package port

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/google/uuid"
)

type EntryGetter interface {
	Get(ctx context.Context, entryID uuid.UUID) (*entity.Entry, error)
}

type EntryOneWithUserGetter interface {
	GetOneWithUser(ctx context.Context, entryID uuid.UUID, user *entity.User) (*entity.Entry, error)
}

type EntryAllByUserGetter interface {
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Entry, error)
}

type EntryAdder interface {
	Add(ctx context.Context, entry *entity.Entry) error
}

type EntryChanger interface {
	Change(ctx context.Context, entry *entity.Entry) error
}

type EntryRemover interface {
	Remove(ctx context.Context, entryID uuid.UUID) error
}

type EntryRepository interface {
	EntryGetter
	EntryOneWithUserGetter
	EntryAllByUserGetter
	EntryAdder
	EntryChanger
	EntryRemover
}
