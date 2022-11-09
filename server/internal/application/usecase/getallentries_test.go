package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllEntries(t *testing.T) {
	type m struct {
		entryRepository *mocks.EntryAllByUserGetter
		userRepository  *mocks.UserGetter
	}

	testUserID := uuid.New()
	testEntries := []*entity.Entry{
		{
			ID:   uuid.New(),
			Name: "entry1",
		},
		{
			ID:   uuid.New(),
			Name: "entry2",
		},
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	tests := []struct {
		name        string
		args        *args
		wantErr     error
		mockPrepare func(a *args) *m
	}{
		{
			"success",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				userID: testUserID,
			},
			nil,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryRepository := mocks.NewEntryAllByUserGetter(t)
				entryRepository.EXPECT().
					GetAllByUser(a.ctx, testUser).
					Return(testEntries, nil)

				return &m{entryRepository, userRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			types, err := GetAllEntries(
				tt.args.ctx,
				m.entryRepository,
				m.userRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testEntries[0].Name, types[0].Name)
		})
	}
}

func TestGetAllEntriesFail(t *testing.T) {
	type m struct {
		entryRepository *mocks.EntryAllByUserGetter
		userRepository  *mocks.UserGetter
	}

	testUserID := uuid.New()
	testEntries := []*entity.Entry{
		{
			ID:   uuid.New(),
			Name: "entry1",
		},
		{
			ID:   uuid.New(),
			Name: "entry2",
		},
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	tests := []struct {
		name        string
		args        *args
		wantErr     error
		mockPrepare func(a *args) *m
	}{
		{
			"fail with incorrect session",
			&args{
				ctx:    context.WithValue(context.Background(), port.SessionContextKey, nil),
				userID: testUserID,
			},
			ErrIncorrectSession,
			func(a *args) *m {
				userRepository := mocks.NewUserGetter(t)
				entryRepository := mocks.NewEntryAllByUserGetter(t)

				return &m{entryRepository, userRepository}
			},
		},
		{
			"fail with user not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				userID: testUserID,
			},
			ErrUserNotFound,
			func(a *args) *m {
				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(nil, ErrUserNotFound)

				entryRepository := mocks.NewEntryAllByUserGetter(t)

				return &m{entryRepository, userRepository}
			},
		},
		{
			"fail with entries not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				userID: testUserID,
			},
			ErrEntryNotFound,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryRepository := mocks.NewEntryAllByUserGetter(t)
				entryRepository.EXPECT().
					GetAllByUser(a.ctx, testUser).
					Return(nil, ErrEntryNotFound)

				return &m{entryRepository, userRepository}
			},
		},
		{
			"fail with entries unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				userID: testUserID,
			},
			ErrInternalServerError,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryRepository := mocks.NewEntryAllByUserGetter(t)
				entryRepository.EXPECT().
					GetAllByUser(a.ctx, testUser).
					Return(nil, fmt.Errorf("unexpected error"))

				return &m{entryRepository, userRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			types, err := GetAllEntries(
				tt.args.ctx,
				m.entryRepository,
				m.userRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testEntries[0].Name, types[0].Name)
		})
	}
}
