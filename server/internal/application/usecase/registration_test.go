//go:build unit

package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegistration(t *testing.T) {
	type m struct {
		userRegistrator *mocks.UserRegistrator
		hasher          *mocks.PasswordHasher
	}

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
				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(nil, ErrUserNotFound)
				userRegistrator.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(nil).
					Once()

				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.regData.Password).
					Return("zzz")

				return &m{userRegistrator, passwordHasher}
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
				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(
						&entity.User{
							ID:           uuid.New(),
							Login:        "other login",
							PasswordHash: "other hash",
							Current:      0,
						},
						nil,
					)

				passwordHasher := mocks.NewPasswordHasher(t)

				return &m{userRegistrator, passwordHasher}
			},
		},
		{
			"fail with other error",
			&args{
				context.Background(),
				RegistrationData{
					Login:    "login",
					Password: "password",
				},
			},
			fmt.Errorf("other error"),
			func(a *args) *m {
				userRegistrator := mocks.NewUserRegistrator(t)
				userRegistrator.EXPECT().
					GetByLogin(a.ctx, a.regData.Login).
					Return(
						nil,
						fmt.Errorf("other error"),
					)

				passwordHasher := mocks.NewPasswordHasher(t)

				return &m{userRegistrator, passwordHasher}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			user, err := Registration(
				tt.args.ctx,
				tt.args.regData,
				m.userRegistrator,
				m.hasher,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.args.regData.Login, user.Login)
		})
	}
}
