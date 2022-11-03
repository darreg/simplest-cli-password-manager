package repository

import (
	"context"
	"database/sql"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
)

type EntryRepository struct {
	tx *adapter.Transactor
	db *sql.DB
}

func NewEntryRepository(tx *adapter.Transactor, db *sql.DB) *EntryRepository {
	return &EntryRepository{tx: tx, db: db}
}

func (e EntryRepository) Get(ctx context.Context, entryID uuid.UUID) (*entity.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (e EntryRepository) GetByUser(ctx context.Context, user *entity.User) ([]*entity.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (e EntryRepository) Add(ctx context.Context, entry *entity.Entry) error {
	//TODO implement me
	panic("implement me")
}

func (e EntryRepository) Change(ctx context.Context, entry *entity.Entry) error {
	//TODO implement me
	panic("implement me")
}

func (e EntryRepository) Remove(ctx context.Context, entryID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

