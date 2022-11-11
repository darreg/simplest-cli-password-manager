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

func TestSelectLoginMethod(t *testing.T) {
	type m struct {
		cliScript *mocks.CLISelectLoginMethodSupporter
	}

	testError := fmt.Errorf("unexpected error")
	loginMethodNames := []string{"name1", "name2"}

	type args struct {
		ctx          context.Context
		loginMethods map[string]func() (string, error)
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
				loginMethods: map[string]func() (string, error){
					loginMethodNames[0]: func() (string, error) { return "name1", nil },
					loginMethodNames[1]: func() (string, error) { return "name2", nil },
				},
			},
			nil,
			func(a *args) *m {

				cliScript := mocks.NewCLISelectLoginMethodSupporter(t)
				cliScript.EXPECT().
					SelectLoginMethod(a.ctx, loginMethodNames).
					Return(loginMethodNames[1], nil)

				return &m{cliScript}
			},
		},
		{
			"fail with unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
				loginMethods: map[string]func() (string, error){
					loginMethodNames[0]: func() (string, error) { return "key1", nil },
					loginMethodNames[1]: func() (string, error) { return "key2", nil },
				},
			},
			testError,
			func(a *args) *m {

				cliScript := mocks.NewCLISelectLoginMethodSupporter(t)
				cliScript.EXPECT().
					SelectLoginMethod(a.ctx, loginMethodNames).
					Return("", testError)

				return &m{cliScript}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			result, err := SelectLoginMethod(
				tt.args.ctx,
				m.cliScript,
				tt.args.loginMethods,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			r1, _ := tt.args.loginMethods[loginMethodNames[1]]()
			r2, _ := result()
			assert.Equal(t, r1, r2)
		})
	}
}
