//go:build unit

package usecase

import (
	"context"
	"fmt"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEntry(t *testing.T) {
	type m struct {
		decryptor       *mocks.Decryptor
		entryRepository *mocks.EntryOneWithUserGetter
		userRepository  *mocks.UserGetter
	}

	testUserID := uuid.New()
	testRequestedEntryUUID := uuid.New()
	testData := []byte("тестовые данные")

	type args struct {
		ctx                context.Context
		requestedEntryUUID uuid.UUID
		requestedEntryID   string
		userID             uuid.UUID
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
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   testRequestedEntryUUID.String(),
				userID:             testUserID,
			},
			nil,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				entryGetter.EXPECT().
					GetOneWithUser(a.ctx, a.requestedEntryUUID, testUser).
					Return(&entity.Entry{
						ID:     a.requestedEntryUUID,
						UserID: testUser.ID,
						Data:   testData,
					}, nil)

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(string(testData)).
					Return(testData, nil)

				return &m{decryptor, entryGetter, userGetter}
			},
		},
		{
			"fail with incorrect session",
			&args{
				ctx:                context.WithValue(context.Background(), port.SessionContextKey, nil),
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   testRequestedEntryUUID.String(),
				userID:             testUserID,
			},
			ErrIncorrectSession,
			func(a *args) *m {

				userGetter := mocks.NewUserGetter(t)
				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				decryptor := mocks.NewDecryptor(t)

				return &m{decryptor, entryGetter, userGetter}
			},
		},
		{
			"fail with user not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   testRequestedEntryUUID.String(),
				userID:             testUserID,
			},
			ErrUserNotFound,
			func(a *args) *m {

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(nil, ErrUserNotFound)
				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				decryptor := mocks.NewDecryptor(t)

				return &m{decryptor, entryGetter, userGetter}
			},
		},
		{
			"fail with entry uuid parse error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   "incorrect",
				userID:             testUserID,
			},
			ErrInvalidArgument,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				decryptor := mocks.NewDecryptor(t)

				return &m{decryptor, entryGetter, userGetter}
			},
		},
		{
			"fail with entry not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   testRequestedEntryUUID.String(),
				userID:             testUserID,
			},
			ErrEntryNotFound,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				entryGetter.EXPECT().
					GetOneWithUser(a.ctx, a.requestedEntryUUID, testUser).
					Return(nil, ErrEntryNotFound)

				decryptor := mocks.NewDecryptor(t)

				return &m{decryptor, entryGetter, userGetter}
			},
		},
		{
			"fail with entry repository unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   testRequestedEntryUUID.String(),
				userID:             testUserID,
			},
			ErrInternalServerError,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				entryGetter.EXPECT().
					GetOneWithUser(a.ctx, a.requestedEntryUUID, testUser).
					Return(nil, fmt.Errorf("unexpected error"))

				decryptor := mocks.NewDecryptor(t)

				return &m{decryptor, entryGetter, userGetter}
			},
		},
		{
			"fail with decryption error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				requestedEntryUUID: testRequestedEntryUUID,
				requestedEntryID:   testRequestedEntryUUID.String(),
				userID:             testUserID,
			},
			ErrInternalServerError,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				entryGetter := mocks.NewEntryOneWithUserGetter(t)
				entryGetter.EXPECT().
					GetOneWithUser(a.ctx, a.requestedEntryUUID, testUser).
					Return(&entity.Entry{
						ID:     a.requestedEntryUUID,
						UserID: testUser.ID,
						Data:   testData,
					}, nil)

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(string(testData)).
					Return(nil, fmt.Errorf("unexpected error"))

				return &m{decryptor, entryGetter, userGetter}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			entry, err := GetEntry(
				tt.args.ctx,
				tt.args.requestedEntryID,
				m.decryptor,
				m.entryRepository,
				m.userRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.args.requestedEntryID, entry.ID.String())
		})
	}
}
