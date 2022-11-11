package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
)

// EntryRepository implements a repository for entry.
type EntryRepository struct {
	tx *adapter.Transactor
}

func NewEntryRepository(tx *adapter.Transactor) *EntryRepository {
	return &EntryRepository{tx: tx}
}

func (e EntryRepository) Get(ctx context.Context, entryID uuid.UUID) (*entity.Entry, error) {
	var entry entity.Entry
	var createdAt time.Time
	var updatedAt time.Time

	err := e.tx.QueryRowContext(ctx,
		"SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at FROM entries WHERE id = $1",
		entryID,
	).Scan(
		&entry.ID,
		&entry.UserID,
		&entry.TypeID,
		&entry.Name,
		&entry.Metadata,
		&entry.Data,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrEntryNotFound
		}
		return nil, err
	}

	entry.CreatedAt = &createdAt
	entry.UpdatedAt = &updatedAt

	return &entry, nil
}

func (e EntryRepository) GetOneWithUser(
	ctx context.Context,
	entryID uuid.UUID,
	user *entity.User,
) (*entity.Entry, error) {
	var entry entity.Entry
	var createdAt time.Time
	var updatedAt time.Time

	err := e.tx.QueryRowContext(ctx,
		"SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at "+
			"FROM entries WHERE id = $1 AND user_id = $2",
		entryID, user.ID,
	).Scan(
		&entry.ID,
		&entry.UserID,
		&entry.TypeID,
		&entry.Name,
		&entry.Metadata,
		&entry.Data,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrEntryNotFound
		}
		return nil, err
	}

	entry.CreatedAt = &createdAt
	entry.UpdatedAt = &updatedAt

	return &entry, nil
}

func (e EntryRepository) GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Entry, error) {
	rows, err := e.tx.QueryContext(ctx,
		"SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at "+
			"FROM entries WHERE user_id = $1",
		user.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrEntryNotFound
		}
		return nil, err
	}

	defer rows.Close()

	entries := make([]*entity.Entry, 0)
	for rows.Next() {
		var entry entity.Entry
		var createdAt time.Time
		var updatedAt time.Time
		err = rows.Scan(
			&entry.ID,
			&entry.UserID,
			&entry.TypeID,
			&entry.Name,
			&entry.Metadata,
			&entry.Data,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		entry.CreatedAt = &createdAt
		entry.UpdatedAt = &updatedAt

		entries = append(entries, &entry)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, usecase.ErrEntryNotFound
	}

	return entries, nil
}

func (e EntryRepository) Add(ctx context.Context, entry *entity.Entry) error {
	_, err := e.tx.ExecContext(ctx,
		"INSERT "+
			"INTO entries (id, user_id, type_id, name, metadata, data, created_at, updated_at) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		entry.ID, entry.UserID, entry.TypeID, entry.Name, entry.Metadata, entry.Data, entry.CreatedAt, entry.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (e EntryRepository) Change(ctx context.Context, entry *entity.Entry) error {
	_, err := e.tx.ExecContext(ctx,
		"UPDATE entries "+
			"SET user_id=$2, type_id=$3, name=$4, metadata=$5, data=$6, updated_at=$7 WHERE id=$1",
		entry.ID, entry.UserID, entry.TypeID, entry.Name, entry.Metadata, entry.Data, entry.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (e EntryRepository) Remove(ctx context.Context, entryID uuid.UUID) error {
	_, err := e.tx.ExecContext(ctx, "DELETE FROM entries WHERE id=$1", entryID)
	return err
}
