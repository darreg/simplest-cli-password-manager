//go:build unit

package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserGet(t *testing.T) {
	testUUID := uuid.New()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		want := &entity.User{
			ID:           uuid.New(),
			Name:         "UserName",
			Login:        "qwerty",
			PasswordHash: "pass",
		}

		mock.
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, login, password FROM users WHERE id = $1")).
			WithArgs(want.ID).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "name", "login", "password"}).
				AddRow(want.ID, want.Name, want.Login, want.PasswordHash),
			)

		repository := NewUserRepository(&adapter.Transactor{DB: db})
		got, err := repository.Get(context.Background(), want.ID)
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, want.ID, got.ID)

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
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, login, password FROM users WHERE id = $1")).
			WithArgs(testUUID).
			WillReturnError(sql.ErrNoRows)

		repository := NewUserRepository(&adapter.Transactor{DB: db})
		_, err = repository.Get(context.Background(), testUUID)
		assert.ErrorIs(t, err, usecase.ErrUserNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUserGetByLogin(t *testing.T) {
	want := &entity.User{
		ID:           uuid.New(),
		Login:        "qwerty",
		PasswordHash: "pass",
	}

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, login, password FROM users WHERE login = $1")).
			WithArgs(want.Login).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "name", "login", "password"}).
				AddRow(want.ID, want.Name, want.Login, want.PasswordHash),
			)

		repository := NewUserRepository(&adapter.Transactor{DB: db})
		got, err := repository.GetByLogin(context.Background(), want.Login)
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, want.ID, got.ID)

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
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, login, password FROM users WHERE login = $1")).
			WithArgs(want.Login).
			WillReturnError(sql.ErrNoRows)

		repository := NewUserRepository(&adapter.Transactor{DB: db})
		_, err = repository.GetByLogin(context.Background(), want.Login)
		assert.ErrorIs(t, err, usecase.ErrUserNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUserGetByCredential(t *testing.T) {
	want := &entity.User{
		ID:           uuid.New(),
		Login:        "qwerty",
		PasswordHash: "pass",
	}

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, login, password FROM users WHERE login = $1 AND password=$2")).
			WithArgs(want.Login, want.PasswordHash).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "name", "login", "password"}).
				AddRow(want.ID, want.Name, want.Login, want.PasswordHash),
			)

		repository := NewUserRepository(&adapter.Transactor{DB: db})
		got, err := repository.GetByCredential(context.Background(), want.Login, want.PasswordHash)
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, want.ID, got.ID)

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
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, login, password FROM users WHERE login = $1 AND password=$2")).
			WithArgs(want.Login, want.PasswordHash).
			WillReturnError(sql.ErrNoRows)

		repository := NewUserRepository(&adapter.Transactor{DB: db})
		_, err = repository.GetByCredential(context.Background(), want.Login, want.PasswordHash)
		assert.ErrorIs(t, err, usecase.ErrUserNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUserAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	arg := &entity.User{
		ID:           uuid.New(),
		Name:         "UserName",
		Login:        "qwerty",
		PasswordHash: "pass",
	}

	mock.
		ExpectExec(regexp.QuoteMeta("INSERT INTO users(id, name, login, password) VALUES($1, $2, $3, $4)")).
		WithArgs(arg.ID, arg.Name, arg.Login, arg.PasswordHash).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewUserRepository(&adapter.Transactor{DB: db})
	err = repository.Add(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserChange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	arg := &entity.User{
		ID:           uuid.New(),
		Login:        "qwerty",
		PasswordHash: "pass",
	}

	mock.
		ExpectExec(regexp.QuoteMeta("UPDATE users SET login=$2 WHERE id=$1")).
		WithArgs(arg.ID, arg.Login).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewUserRepository(&adapter.Transactor{DB: db})
	err = repository.Change(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	arg := &entity.User{
		ID:           uuid.New(),
		Login:        "qwerty",
		PasswordHash: "pass",
	}

	mock.
		ExpectExec(regexp.QuoteMeta("UPDATE users SET password=$2 WHERE id=$1")).
		WithArgs(arg.ID, arg.PasswordHash).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewUserRepository(&adapter.Transactor{DB: db})
	err = repository.ChangePassword(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.
		ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id=$1")).
		WithArgs(testUUID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewUserRepository(&adapter.Transactor{DB: db})
	err = repository.Remove(context.Background(), testUUID)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
