//go:build unit

package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionGet(t *testing.T) {
	testUUID := uuid.New()
	userUUID := uuid.New()
	testTime := time.Now()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		want := &entity.Session{
			ID:           testUUID,
			UserID:       userUUID,
			LoginTime:    &testTime,
			LastSeenTime: &testTime,
		}

		mock.
			ExpectQuery(regexp.QuoteMeta(
				"SELECT id, user_id, login_time, last_seen_time FROM sessions WHERE id = $1",
			)).
			WithArgs(want.ID).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "user_id", "login_time", "last_seen_time"}).
				AddRow(want.ID, want.UserID, want.LoginTime, want.LastSeenTime),
			)

		repository := NewSessionRepository(&adapter.Transactor{DB: db})
		got, err := repository.Get(context.Background(), want.ID)
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, want.LoginTime, got.LoginTime)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(regexp.QuoteMeta(
				"SELECT id, user_id, login_time, last_seen_time FROM sessions WHERE id = $1",
			)).
			WithArgs(testUUID).
			WillReturnError(sql.ErrNoRows)

		repository := NewSessionRepository(&adapter.Transactor{DB: db})
		_, err = repository.Get(context.Background(), testUUID)
		require.NotNil(t, err)
		assert.ErrorIs(t, err, usecase.ErrSessionNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

}

func TestSessionAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()
	userUUID := uuid.New()
	testTime := time.Now()

	arg := &entity.Session{
		ID:           testUUID,
		UserID:       userUUID,
		LoginTime:    &testTime,
		LastSeenTime: &testTime,
	}

	mock.
		ExpectExec(regexp.QuoteMeta(
			"INSERT INTO sessions (id, user_id, login_time, last_seen_time) VALUES ($1, $2, $3, $4)",
		)).
		WithArgs(arg.ID, arg.UserID, arg.LoginTime, arg.LastSeenTime).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewSessionRepository(&adapter.Transactor{DB: db})
	err = repository.Add(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSessionChange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()
	userUUID := uuid.New()
	testTime := time.Now()

	arg := &entity.Session{
		ID:           testUUID,
		UserID:       userUUID,
		LoginTime:    &testTime,
		LastSeenTime: &testTime,
	}

	mock.
		ExpectExec(regexp.QuoteMeta("UPDATE sessions SET last_seen_time=$2 WHERE id=$1")).
		WithArgs(arg.ID, arg.LastSeenTime).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewSessionRepository(&adapter.Transactor{DB: db})
	err = repository.Change(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSessionRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.
		ExpectExec(regexp.QuoteMeta("DELETE FROM sessions WHERE id=$1")).
		WithArgs(testUUID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewSessionRepository(&adapter.Transactor{DB: db})
	err = repository.Remove(context.Background(), testUUID)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
