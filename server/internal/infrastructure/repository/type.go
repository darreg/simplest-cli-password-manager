package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
)

// TypeRepository implements a repository for type.
type TypeRepository struct {
	tx *adapter.Transactor
}

func NewTypeRepository(tx *adapter.Transactor) *TypeRepository {
	return &TypeRepository{tx: tx}
}

func (t TypeRepository) Get(ctx context.Context, tpID uuid.UUID) (*entity.Type, error) {
	var tp entity.Type

	err := t.tx.QueryRowContext(ctx,
		"SELECT id, name, is_binary FROM types WHERE id = $1", tpID,
	).Scan(&tp.ID, &tp.Name, &tp.IsBinary)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrTypeNotFound
		}
		return nil, err
	}

	return &tp, nil
}

func (t TypeRepository) GetAll(ctx context.Context) ([]*entity.Type, error) {
	rows, err := t.tx.QueryContext(ctx,
		"SELECT id, name, is_binary FROM types",
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrTypeNotFound
		}
		return nil, err
	}

	defer rows.Close()

	types := make([]*entity.Type, 0)
	for rows.Next() {
		var tp entity.Type
		err = rows.Scan(
			&tp.ID,
			&tp.Name,
			&tp.IsBinary,
		)
		if err != nil {
			return nil, err
		}

		types = append(types, &tp)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(types) == 0 {
		return nil, usecase.ErrTypeNotFound
	}

	return types, nil
}

func (t TypeRepository) Add(ctx context.Context, tp *entity.Type) error {
	_, err := t.tx.ExecContext(ctx,
		"INSERT INTO types (id, name, is_binary) VALUES ($1, $2, $3)",
		tp.ID, tp.Name, tp.IsBinary)
	if err != nil {
		return err
	}

	return nil
}

func (t TypeRepository) Change(ctx context.Context, tp *entity.Type) error {
	_, err := t.tx.ExecContext(ctx,
		"UPDATE types SET name=$2, is_binary=$3 WHERE id=$1",
		tp.ID, tp.Name, tp.IsBinary)
	if err != nil {
		return err
	}

	return nil
}

func (t TypeRepository) Remove(ctx context.Context, tpID uuid.UUID) error {
	_, err := t.tx.ExecContext(ctx, "DELETE FROM types WHERE id=$1", tpID)
	return err
}
