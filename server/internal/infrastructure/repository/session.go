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

// SessionRepository implements a repository for session.
type SessionRepository struct {
	tx *adapter.Transactor
}

func NewSessionRepository(tx *adapter.Transactor) *SessionRepository {
	return &SessionRepository{tx: tx}
}

func (t SessionRepository) Get(ctx context.Context, sessionID uuid.UUID) (*entity.Session, error) {
	var session entity.Session

	err := t.tx.QueryRowContext(ctx,
		"SELECT id, user_id, login_time, last_seen_time FROM sessions WHERE id = $1", sessionID,
	).Scan(&session.ID, &session.UserID, &session.LoginTime, &session.LastSeenTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrSessionNotFound
		}
		return nil, err
	}

	return &session, nil
}

func (t SessionRepository) Add(ctx context.Context, session *entity.Session) error {
	_, err := t.tx.ExecContext(ctx,
		"INSERT INTO sessions (id, user_id, login_time, last_seen_time) VALUES ($1, $2, $3, $4)",
		session.ID, session.UserID, session.LoginTime, session.LastSeenTime)
	if err != nil {
		return err
	}

	return nil
}

func (t SessionRepository) Change(ctx context.Context, session *entity.Session) error {
	_, err := t.tx.ExecContext(ctx,
		"UPDATE sessions SET last_seen_time=$2 WHERE id=$1",
		session.ID, session.LastSeenTime)
	if err != nil {
		return err
	}

	return nil
}

func (t SessionRepository) Remove(ctx context.Context, sessionID uuid.UUID) error {
	_, err := t.tx.ExecContext(ctx, "DELETE FROM sessions WHERE id=$1", sessionID)
	return err
}
