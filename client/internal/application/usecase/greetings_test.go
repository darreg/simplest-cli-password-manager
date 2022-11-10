//go:build unit

package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGreetings(t *testing.T) {
	type m struct {
		userRepository *mocks.GRPCClientUserGetter
	}

	testError := fmt.Errorf("unexpected error")
	testUser := &model.User{
		ID:    "xxx-yyy-zzz",
		Name:  "User Name",
		Login: "User Login",
	}

	type args struct {
		ctx context.Context
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
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			nil,
			func(a *args) *m {

				userRepository := mocks.NewGRPCClientUserGetter(t)
				userRepository.EXPECT().
					GetUser(a.ctx).
					Return(testUser, nil)

				return &m{userRepository}
			},
		},
		{
			"fail with unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {

				userRepository := mocks.NewGRPCClientUserGetter(t)
				userRepository.EXPECT().
					GetUser(a.ctx).
					Return(nil, testError)

				return &m{userRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			result, err := Greetings(
				tt.args.ctx,
				m.userRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("Greetings, %s!", testUser.Name), result)
		})
	}
}
