//go:build unit

package usecase

import (
	"context"
	"fmt"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLogin(t *testing.T) {
	type m struct {
		hasher            *mocks.PasswordHasher
		encryptor         *mocks.Encryptor
		userRepository    *mocks.UserByCredentialGetter
		sessionRepository *mocks.SessionAdder
	}

	testUserID := uuid.New()
	testEncryptedSessionKey := "encrypted"

	type args struct {
		ctx  context.Context
		cred Credential
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
				context.Background(),
				Credential{
					Login:    "login",
					Password: "password",
				},
			},
			nil,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				userRepository := mocks.NewUserByCredentialGetter(t)
				userRepository.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(
						&entity.User{
							ID:           testUserID,
							Login:        a.cred.Login,
							PasswordHash: a.cred.Password,
						},
						nil,
					)

				sessionRepository := mocks.NewSessionAdder(t)
				sessionRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Session")).
					Return(nil).
					Once()

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return(testEncryptedSessionKey, nil)

				return &m{passwordHasher, encryptor, userRepository, sessionRepository}
			},
		},
		{
			"fail with encryptor unexpected error",
			&args{
				context.Background(),
				Credential{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				userRepository := mocks.NewUserByCredentialGetter(t)
				userRepository.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(
						&entity.User{
							ID:           testUserID,
							Login:        a.cred.Login,
							PasswordHash: a.cred.Password,
						},
						nil,
					)

				sessionRepository := mocks.NewSessionAdder(t)
				sessionRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Session")).
					Return(nil).
					Once()

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return("", fmt.Errorf("unexpected error"))

				return &m{passwordHasher, encryptor, userRepository, sessionRepository}
			},
		},
		{
			"fail with user not found",
			&args{
				context.Background(),
				Credential{
					Login:    "login",
					Password: "password",
				},
			},
			ErrNotAuthenticated,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				userRepository := mocks.NewUserByCredentialGetter(t)
				userRepository.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(nil, ErrUserNotFound)

				sessionRepository := mocks.NewSessionAdder(t)

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRepository, sessionRepository}
			},
		},
		{
			"fail with user repository unexpected error",
			&args{
				context.Background(),
				Credential{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				userRepository := mocks.NewUserByCredentialGetter(t)
				userRepository.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(nil, fmt.Errorf("test error"))

				sessionRepository := mocks.NewSessionAdder(t)

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRepository, sessionRepository}
			},
		},
		{
			"fail with session repository unexpected error",
			&args{
				context.Background(),
				Credential{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				userRepository := mocks.NewUserByCredentialGetter(t)
				userRepository.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(
						&entity.User{
							ID:           testUserID,
							Login:        a.cred.Login,
							PasswordHash: a.cred.Password,
						},
						nil,
					)

				sessionRepository := mocks.NewSessionAdder(t)
				sessionRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Session")).
					Return(fmt.Errorf("test error"))

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRepository, sessionRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			encryptedSessionKey, err := Login(
				tt.args.ctx,
				tt.args.cred,
				m.hasher,
				m.encryptor,
				m.userRepository,
				m.sessionRepository,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testEncryptedSessionKey, encryptedSessionKey)
		})
	}
}
