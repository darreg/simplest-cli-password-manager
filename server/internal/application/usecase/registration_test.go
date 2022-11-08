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

func TestRegistration(t *testing.T) {
	type m struct {
		hasher            *mocks.PasswordHasher
		encryptor         *mocks.Encryptor
		userRegistrator   *mocks.UserRegistrator
		sessionRepository *mocks.SessionAdder
	}

	testEncryptedSessionKey := "encrypted"

	type args struct {
		ctx     context.Context
		regData RegistrationData
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
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			nil,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.regData.Password).
					Return("zzz")

				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(nil, ErrUserNotFound)
				userRegistrator.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(nil).
					Once()

				sessionRepository := mocks.NewSessionAdder(t)
				sessionRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Session")).
					Return(nil).
					Once()

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return(testEncryptedSessionKey, nil)

				return &m{passwordHasher, encryptor, userRegistrator, sessionRepository}
			},
		},
		{
			"fail with encryptor unexpected error",
			&args{
				context.Background(),
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.regData.Password).
					Return("zzz")

				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(nil, ErrUserNotFound)
				userRegistrator.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(nil).
					Once()

				sessionRepository := mocks.NewSessionAdder(t)
				sessionRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Session")).
					Return(nil).
					Once()

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return("", fmt.Errorf("unexpected error"))

				return &m{passwordHasher, encryptor, userRegistrator, sessionRepository}
			},
		},
		{
			"fail with user repository get unexpected error",
			&args{
				context.Background(),
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)

				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(nil, fmt.Errorf("test error"))

				sessionRepository := mocks.NewSessionAdder(t)

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRegistrator, sessionRepository}
			},
		},
		{
			"fail with login already use",
			&args{
				context.Background(),
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			ErrLoginAlreadyUse,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)

				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(&entity.User{
						ID: uuid.New(),
					}, nil)

				sessionRepository := mocks.NewSessionAdder(t)

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRegistrator, sessionRepository}
			},
		},
		{
			"fail with user repository add unexpected error",
			&args{
				context.Background(),
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.regData.Password).
					Return("zzz")

				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(nil, ErrUserNotFound)
				userRegistrator.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(fmt.Errorf("test error")).
					Once()

				sessionRepository := mocks.NewSessionAdder(t)

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRegistrator, sessionRepository}
			},
		},
		{
			"fail with session repository unexpected error",
			&args{
				context.Background(),
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			ErrInternalServerError,
			func(a *args) *m {
				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.regData.Password).
					Return("zzz")

				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(nil, ErrUserNotFound)
				userRegistrator.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(nil).
					Once()

				sessionRepository := mocks.NewSessionAdder(t)
				sessionRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Session")).
					Return(fmt.Errorf("test error")).
					Once()

				encryptor := mocks.NewEncryptor(t)

				return &m{passwordHasher, encryptor, userRegistrator, sessionRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			session, err := Registration(
				tt.args.ctx,
				tt.args.regData,
				m.hasher,
				m.encryptor,
				m.userRegistrator,
				m.sessionRepository,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, session)
		})
	}
}
