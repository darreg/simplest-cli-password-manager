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

// UserRepository implements a repository for user.
type UserRepository struct {
	tx *adapter.Transactor
}

func NewUserRepository(tx *adapter.Transactor) *UserRepository {
	return &UserRepository{tx: tx}
}

func (u UserRepository) Get(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := u.tx.QueryRowContext(ctx,
		"SELECT id, name, login, password FROM users WHERE id = $1", userID,
	).Scan(&user.ID, &user.Name, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	var user entity.User

	err := u.tx.QueryRowContext(ctx,
		"SELECT id, name, login, password FROM users WHERE login = $1", login,
	).Scan(&user.ID, &user.Name, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetByCredential(ctx context.Context, login, passwordHash string) (*entity.User, error) {
	var user entity.User

	err := u.tx.QueryRowContext(ctx,
		"SELECT id, name, login, password FROM users WHERE login = $1 AND password=$2", login, passwordHash,
	).Scan(&user.ID, &user.Name, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) Add(ctx context.Context, user *entity.User) error {
	_, err := u.tx.ExecContext(ctx,
		"INSERT INTO users(id, name, login, password) VALUES($1, $2, $3, $4)",
		user.ID, user.Name, user.Login, user.PasswordHash,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Change(ctx context.Context, user *entity.User) error {
	_, err := u.tx.ExecContext(ctx,
		"UPDATE users SET login=$2 WHERE id=$1", user.ID, user.Login,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) ChangePassword(ctx context.Context, user *entity.User) error {
	_, err := u.tx.ExecContext(ctx, "UPDATE users SET password=$2 WHERE id=$1", user.ID, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Remove(ctx context.Context, userID uuid.UUID) error {
	_, err := u.tx.ExecContext(ctx, "DELETE FROM users WHERE id=$1", userID)
	return err
}
