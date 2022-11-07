//go:build unit

package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSessionValidate(t *testing.T) {
	type m struct {
		decryptor         *mocks.Decryptor
		sessionRepository *mocks.SessionGetter
	}

	type args struct {
		ctx                 context.Context
		encryptedSessionKey string
		sessionLifeTime     string
		userID              uuid.UUID
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
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			nil,
			func(a *args) *m {
				sessionID := uuid.New()
				testTime := time.Now().Add(-time.Minute * 59)

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte(sessionID.String()), nil)

				sessionGetter := mocks.NewSessionGetter(t)
				sessionGetter.EXPECT().
					Get(a.ctx, sessionID).
					Return(&entity.Session{
						ID:           sessionID,
						UserID:       a.userID,
						LoginTime:    &testTime,
						LastSeenTime: &testTime,
					}, nil)

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail with expired session",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			ErrNotAuthenticated,
			func(a *args) *m {
				sessionID := uuid.New()
				testTime := time.Now().Add(-time.Minute * 61)

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte(sessionID.String()), nil)

				sessionGetter := mocks.NewSessionGetter(t)
				sessionGetter.EXPECT().
					Get(a.ctx, sessionID).
					Return(&entity.Session{
						ID:           sessionID,
						UserID:       a.userID,
						LoginTime:    &testTime,
						LastSeenTime: &testTime,
					}, nil)

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail with empty encryptedSessionKey",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			ErrInvalidSessionKey,
			func(a *args) *m {
				decryptor := mocks.NewDecryptor(t)
				sessionGetter := mocks.NewSessionGetter(t)

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail with decrypt error",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			ErrInvalidSessionKey,
			func(a *args) *m {
				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte{}, errors.New("fake error"))

				sessionGetter := mocks.NewSessionGetter(t)

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail with incorrect userID",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			ErrInternalServerError,
			func(a *args) *m {
				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte("incorrect"), nil)

				sessionGetter := mocks.NewSessionGetter(t)

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail with not found session",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			ErrNotAuthenticated,
			func(a *args) *m {
				sessionID := uuid.New()

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte(sessionID.String()), nil)

				sessionGetter := mocks.NewSessionGetter(t)
				sessionGetter.EXPECT().
					Get(a.ctx, sessionID).
					Return(nil, ErrSessionNotFound)

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail with get error",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "1h",
				userID:              uuid.New(),
			},
			ErrInternalServerError,
			func(a *args) *m {
				sessionID := uuid.New()

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte(sessionID.String()), nil)

				sessionGetter := mocks.NewSessionGetter(t)
				sessionGetter.EXPECT().
					Get(a.ctx, sessionID).
					Return(nil, errors.New("fake error"))

				return &m{decryptor, sessionGetter}
			},
		},
		{
			"fail error parse sessionLifeTime",
			&args{
				ctx:                 context.Background(),
				encryptedSessionKey: "encrypted",
				sessionLifeTime:     "incorrect",
				userID:              uuid.New(),
			},
			ErrInternalServerError,
			func(a *args) *m {
				sessionID := uuid.New()
				testTime := time.Now()

				decryptor := mocks.NewDecryptor(t)
				decryptor.EXPECT().
					Decrypt(a.encryptedSessionKey).
					Return([]byte(sessionID.String()), nil)

				sessionGetter := mocks.NewSessionGetter(t)
				sessionGetter.EXPECT().
					Get(a.ctx, sessionID).
					Return(&entity.Session{
						ID:           sessionID,
						UserID:       a.userID,
						LoginTime:    &testTime,
						LastSeenTime: &testTime,
					}, nil)

				return &m{decryptor, sessionGetter}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			session, err := SessionValidate(
				tt.args.ctx,
				tt.args.encryptedSessionKey,
				tt.args.sessionLifeTime,
				m.decryptor,
				m.sessionRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.args.userID.String(), session.UserID.String())
		})
	}
}
