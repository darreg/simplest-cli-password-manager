package repository

import (
	"context"

	"database/sql"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
)

type TypeRepository struct {
	tx *adapter.Transactor
	db *sql.DB
}

func NewTypeRepository(tx *adapter.Transactor, db *sql.DB) *TypeRepository {
	return &TypeRepository{tx: tx, db: db}
}

func (t TypeRepository) Get(ctx context.Context, tpID uuid.UUID) (*entity.Type, error) {
	//TODO implement me
	panic("implement me")
}

func (t TypeRepository) Add(ctx context.Context, tp *entity.Type) error {
	//TODO implement me
	panic("implement me")
}

func (t TypeRepository) Change(ctx context.Context, tp *entity.Type) error {
	//TODO implement me
	panic("implement me")
}

func (t TypeRepository) Remove(ctx context.Context, tpID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

