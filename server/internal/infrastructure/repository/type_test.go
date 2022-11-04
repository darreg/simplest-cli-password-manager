package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alrund/yp-2-project/server/internal/application/usecase"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	rows := sqlmock.
		NewRows([]string{"id", "name"}).
		AddRow(testUUID, "qwerty")
	query := mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM types WHERE id = $1"))
	query.WithArgs(testUUID).WillReturnRows(rows)

	type args struct {
		ID uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want *entity.Type
	}{
		{
			"success",
			args{
				ID: testUUID,
			},
			&entity.Type{
				ID:   testUUID,
				Name: "qwerty",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTypeRepository(&adapter.Transactor{Db: db})
			got, err := repository.Get(context.Background(), tt.args.ID)
			require.Nil(t, err)
			if tt.want != nil {
				require.NotNil(t, got)
				assert.Equal(t, tt.want.Name, got.Name)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	query := mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM types WHERE id = $1"))
	query.WithArgs(testUUID).WillReturnError(sql.ErrNoRows)

	type args struct {
		ID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Type
		wantErr error
	}{
		{
			"fail",
			args{
				ID: testUUID,
			},
			nil,
			usecase.ErrTypeNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTypeRepository(&adapter.Transactor{Db: db})
			got, err := repository.Get(context.Background(), tt.args.ID)
			if tt.want != nil {
				require.NotNil(t, got)
				assert.Equal(t, tt.want.Name, got.Name)
			}

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO types (id, name) VALUES ($1, $2)")).
		WithArgs(testUUID, "qwerty").
		WillReturnResult(sqlmock.NewResult(0, 1))

	type args struct {
		tp *entity.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				tp: &entity.Type{
					ID:   testUUID,
					Name: "qwerty",
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTypeRepository(&adapter.Transactor{Db: db})
			err := repository.Add(context.Background(), tt.args.tp)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestChange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE types SET name=$2 WHERE id=$1")).
		WithArgs(testUUID, "qwerty").
		WillReturnResult(sqlmock.NewResult(0, 1))

	type args struct {
		tp *entity.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				tp: &entity.Type{
					ID:   testUUID,
					Name: "qwerty",
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTypeRepository(&adapter.Transactor{Db: db})
			err := repository.Change(context.Background(), tt.args.tp)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testUUID := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM types WHERE id=$1")).
		WithArgs(testUUID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	type args struct {
		ID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				ID: testUUID,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := NewTypeRepository(&adapter.Transactor{Db: db})
			err := repository.Remove(context.Background(), tt.args.ID)
			if tt.wantErr {
				assert.NotNil(t, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
