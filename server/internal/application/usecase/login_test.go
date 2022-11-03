//go:build unit

package usecase

import (
	"context"
	"testing"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	type m struct {
		userGetter *mocks.UserByCredentialGetter
		hasher     *mocks.PasswordHasher
	}

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
				userGetter := mocks.NewUserByCredentialGetter(t)
				userGetter.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(
						&entity.User{
							ID:           uuid.New(),
							Login:        a.cred.Login,
							PasswordHash: a.cred.Password,
							Current:      0,
						},
						nil,
					)

				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				return &m{userGetter, passwordHasher}
			},
		},

		{
			"fail with not found",
			&args{
				context.Background(),
				Credential{
					Login:    "login",
					Password: "password",
				},
			},
			ErrUserNotFound,
			func(a *args) *m {
				userGetter := mocks.NewUserByCredentialGetter(t)
				userGetter.EXPECT().
					GetByCredential(a.ctx, a.cred.Login, a.cred.Password).
					Return(nil, ErrUserNotFound)

				passwordHasher := mocks.NewPasswordHasher(t)
				passwordHasher.EXPECT().
					Hash(a.cred.Password).
					Return(a.cred.Password)

				return &m{userGetter, passwordHasher}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			user, err := Login(
				tt.args.ctx,
				tt.args.cred,
				m.userGetter,
				m.hasher,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.args.cred.Login, user.Login)
		})
	}
}
