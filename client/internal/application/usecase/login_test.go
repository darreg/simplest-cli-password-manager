//go:build unit

package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	type m struct {
		client    *mocks.GRPCClientLoginSupporter
		cliScript *mocks.CLILoginSupporter
	}

	testDto := &LoginDTO{
		Login:    "User Login",
		Password: "UserPassword",
	}
	testSessionKey := "test session key"
	testError := fmt.Errorf("unexpected error")

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

				cliScript := mocks.NewCLILoginSupporter(t)
				cliScript.EXPECT().
					Login(a.ctx).
					Return(testDto, nil)

				client := mocks.NewGRPCClientLoginSupporter(t)
				client.EXPECT().
					Login(a.ctx, testDto.Login, testDto.Password).
					Return(testSessionKey, nil)

				return &m{client, cliScript}
			},
		},
		{
			"fail with cliScript.Login unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {

				cliScript := mocks.NewCLILoginSupporter(t)
				cliScript.EXPECT().
					Login(a.ctx).
					Return(nil, testError)

				client := mocks.NewGRPCClientLoginSupporter(t)

				return &m{client, cliScript}
			},
		},
		{
			"fail with internal error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			ErrInternalError,
			func(a *args) *m {

				cliScript := mocks.NewCLILoginSupporter(t)
				cliScript.EXPECT().
					Login(a.ctx).
					Return(&struct{}{}, nil)

				client := mocks.NewGRPCClientLoginSupporter(t)

				return &m{client, cliScript}
			},
		},
		{
			"fail with client.Login unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {

				cliScript := mocks.NewCLILoginSupporter(t)
				cliScript.EXPECT().
					Login(a.ctx).
					Return(testDto, nil)

				client := mocks.NewGRPCClientLoginSupporter(t)
				client.EXPECT().
					Login(a.ctx, testDto.Login, testDto.Password).
					Return("", testError)

				return &m{client, cliScript}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			result, err := Login(
				tt.args.ctx,
				m.client,
				m.cliScript,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testSessionKey, result)
		})
	}
}
