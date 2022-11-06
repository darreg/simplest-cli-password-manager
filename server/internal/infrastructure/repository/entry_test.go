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

func TestEntryGet(t *testing.T) {
	testTime := time.Now()

	want := &entity.Entry{
		ID:        uuid.New(),
		UserID:    uuid.New(),
		TypeID:    uuid.New(),
		Name:      "qwerty",
		Metadata:  "test metadata",
		Data:      []byte("test data"),
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
	}

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at " +
					"FROM entries WHERE id = $1"),
			).
			WithArgs(want.ID).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "user_id", "type_id", "name", "metadata", "data", "created_at", "updated_at"}).
				AddRow(
					want.ID,
					want.UserID,
					want.TypeID,
					want.Name,
					want.Metadata,
					want.Data,
					want.CreatedAt,
					want.UpdatedAt,
				),
			)

		repository := NewEntryRepository(&adapter.Transactor{DB: db})
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
			ExpectQuery(
				regexp.QuoteMeta("SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at " +
					"FROM entries WHERE id = $1"),
			).
			WithArgs(want.ID).
			WillReturnError(sql.ErrNoRows)

		repository := NewEntryRepository(&adapter.Transactor{DB: db})
		_, err = repository.Get(context.Background(), want.ID)
		assert.ErrorIs(t, err, usecase.ErrEntryNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}

func TestEntryGetOneWithUser(t *testing.T) {
	testTime := time.Now()
	testUser := &entity.User{
		ID: uuid.New(),
	}

	want := &entity.Entry{
		ID:        uuid.New(),
		UserID:    testUser.ID,
		TypeID:    uuid.New(),
		Name:      "qwerty",
		Metadata:  "test metadata",
		Data:      []byte("test data"),
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
	}

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at "+
					"FROM entries WHERE id = $1 AND user_id = $2"),
			).
			WithArgs(want.ID, want.UserID).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "user_id", "type_id", "name", "metadata", "data", "created_at", "updated_at"}).
				AddRow(
					want.ID,
					want.UserID,
					want.TypeID,
					want.Name,
					want.Metadata,
					want.Data,
					want.CreatedAt,
					want.UpdatedAt,
				),
			)

		repository := NewEntryRepository(&adapter.Transactor{DB: db})
		got, err := repository.GetOneWithUser(context.Background(), want.ID, testUser)
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
			ExpectQuery(
				regexp.QuoteMeta("SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at "+
					"FROM entries WHERE id = $1 AND user_id = $2"),
			).
			WithArgs(want.ID, want.UserID).
			WillReturnError(sql.ErrNoRows)

		repository := NewEntryRepository(&adapter.Transactor{DB: db})
		_, err = repository.GetOneWithUser(context.Background(), want.ID, testUser)
		assert.ErrorIs(t, err, usecase.ErrEntryNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}

func TestEntryGetAllByUser(t *testing.T) {
	userUUID := uuid.New()
	testTime := time.Now()

	arg := &entity.User{ID: userUUID}
	want := []*entity.Entry{
		{
			ID:        uuid.New(),
			UserID:    userUUID,
			TypeID:    uuid.New(),
			Name:      "qwerty",
			Metadata:  "test metadata",
			Data:      []byte("test data"),
			CreatedAt: &testTime,
			UpdatedAt: &testTime,
		},
		{
			ID:        uuid.New(),
			UserID:    userUUID,
			TypeID:    uuid.New(),
			Name:      "qwerty2",
			Metadata:  "test metadata2",
			Data:      []byte("test data2"),
			CreatedAt: &testTime,
			UpdatedAt: &testTime,
		},
	}

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at " +
					"FROM entries WHERE user_id = $1"),
			).
			WithArgs(userUUID).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "user_id", "type_id", "name", "metadata", "data", "created_at", "updated_at"}).
				AddRow(
					want[0].ID,
					want[0].UserID,
					want[0].TypeID,
					want[0].Name,
					want[0].Metadata,
					want[0].Data,
					want[0].CreatedAt,
					want[0].UpdatedAt,
				).
				AddRow(
					want[1].ID,
					want[1].UserID,
					want[1].TypeID,
					want[1].Name,
					want[1].Metadata,
					want[1].Data,
					want[1].CreatedAt,
					want[1].UpdatedAt,
				),
			)

		repository := NewEntryRepository(&adapter.Transactor{DB: db})
		got, err := repository.GetAllByUser(context.Background(), arg)
		require.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, len(want), len(got))
		assert.Equal(t, want[1].Name, got[1].Name)

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
			ExpectQuery(
				regexp.QuoteMeta("SELECT id, user_id, type_id, name, metadata, data, created_at, updated_at " +
					"FROM entries WHERE user_id = $1"),
			).
			WithArgs(userUUID).
			WillReturnError(sql.ErrNoRows)

		repository := NewEntryRepository(&adapter.Transactor{DB: db})
		_, err = repository.GetAllByUser(context.Background(), arg)
		assert.ErrorIs(t, err, usecase.ErrEntryNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

	})
}

func TestEntryAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testTime := time.Now()

	arg := &entity.Entry{
		ID:        uuid.New(),
		UserID:    uuid.New(),
		TypeID:    uuid.New(),
		Name:      "qwerty",
		Metadata:  "test metadata",
		Data:      []byte("test data"),
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
	}

	mock.
		ExpectExec(regexp.QuoteMeta("INSERT "+
			"INTO entries (id, user_id, type_id, name, metadata, data, created_at, updated_at) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")).
		WithArgs(arg.ID, arg.UserID, arg.TypeID, arg.Name, arg.Metadata, arg.Data, arg.CreatedAt, arg.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewEntryRepository(&adapter.Transactor{DB: db})
	err = repository.Add(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEntryChange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testTime := time.Now()

	arg := &entity.Entry{
		ID:        uuid.New(),
		UserID:    uuid.New(),
		TypeID:    uuid.New(),
		Name:      "qwerty",
		Metadata:  "test metadata",
		Data:      []byte("test data"),
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
	}

	mock.
		ExpectExec(regexp.QuoteMeta("UPDATE entries "+
			"SET user_id=$2, type_id=$3, name=$4, metadata=$5, data=$6, updated_at=$7 WHERE id=$1")).
		WithArgs(arg.ID, arg.UserID, arg.TypeID, arg.Name, arg.Metadata, arg.Data, arg.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewEntryRepository(&adapter.Transactor{DB: db})
	err = repository.Change(context.Background(), arg)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEntryRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.
		ExpectExec(regexp.QuoteMeta("DELETE FROM entries WHERE id=$1")).
		WithArgs(testUUID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewEntryRepository(&adapter.Transactor{DB: db})
	err = repository.Remove(context.Background(), testUUID)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
