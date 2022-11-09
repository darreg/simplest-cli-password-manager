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

func TestTypeGet(t *testing.T) {
	testUUID := uuid.New()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		want := &entity.Type{
			ID:   testUUID,
			Name: "qwerty",
		}

		mock.
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_binary FROM types WHERE id = $1")).
			WithArgs(want.ID).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "name", "is_binary"}).
				AddRow(want.ID, "qwerty", false),
			)

		repository := NewTypeRepository(&adapter.Transactor{DB: db})
		got, err := repository.Get(context.Background(), want.ID)
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, want.Name, got.Name)

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
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_binary FROM types WHERE id = $1")).
			WithArgs(testUUID).
			WillReturnError(sql.ErrNoRows)

		repository := NewTypeRepository(&adapter.Transactor{DB: db})
		_, err = repository.Get(context.Background(), testUUID)
		require.NotNil(t, err)
		assert.ErrorIs(t, err, usecase.ErrTypeNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

}

func TestTypeGetAll(t *testing.T) {
	want := []*entity.Type{
		{
			ID:       uuid.New(),
			Name:     "qwerty",
			IsBinary: false,
		},
		{
			ID:       uuid.New(),
			Name:     "qwerty2",
			IsBinary: false,
		},
	}

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_binary FROM types")).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "name", "is_binary"}).
				AddRow(want[0].ID, want[0].Name, want[0].IsBinary).
				AddRow(want[1].ID, want[1].Name, want[1].IsBinary),
			)

		repository := NewTypeRepository(&adapter.Transactor{DB: db})
		got, err := repository.GetAll(context.Background())
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, want[0].Name, got[0].Name)

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
			ExpectQuery(regexp.QuoteMeta("SELECT id, name, is_binary FROM types")).
			WillReturnError(sql.ErrNoRows)

		repository := NewTypeRepository(&adapter.Transactor{DB: db})
		_, err = repository.GetAll(context.Background())
		require.NotNil(t, err)
		assert.ErrorIs(t, err, usecase.ErrTypeNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestTypeAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	arg := &entity.Type{
		ID:   uuid.New(),
		Name: "qwerty",
	}

	mock.
		ExpectExec(regexp.QuoteMeta("INSERT INTO types (id, name, is_binary) VALUES ($1, $2, $3)")).
		WithArgs(arg.ID, arg.Name, arg.IsBinary).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewTypeRepository(&adapter.Transactor{DB: db})
	err = repository.Add(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTypeChange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	arg := &entity.Type{
		ID:   uuid.New(),
		Name: "qwerty",
	}

	mock.
		ExpectExec(regexp.QuoteMeta("UPDATE types SET name=$2, is_binary=$3 WHERE id=$1")).
		WithArgs(arg.ID, arg.Name, arg.IsBinary).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewTypeRepository(&adapter.Transactor{DB: db})
	err = repository.Change(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTypeRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.
		ExpectExec(regexp.QuoteMeta("DELETE FROM types WHERE id=$1")).
		WithArgs(testUUID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewTypeRepository(&adapter.Transactor{DB: db})
	err = repository.Remove(context.Background(), testUUID)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
