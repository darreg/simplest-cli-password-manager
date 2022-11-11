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

func TestSelectCommand(t *testing.T) {
	type m struct {
		cliScript *mocks.CLISelectCommandSupporter
	}

	testError := fmt.Errorf("unexpected error")
	commandNames := []string{"name1", "name2"}

	type args struct {
		ctx      context.Context
		commands map[string]func() (string, error)
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
				commands: map[string]func() (string, error){
					commandNames[0]: func() (string, error) { return "key1", nil },
					commandNames[1]: func() (string, error) { return "key2", nil },
				},
			},
			nil,
			func(a *args) *m {

				cliScript := mocks.NewCLISelectCommandSupporter(t)
				cliScript.EXPECT().
					SelectCommand(a.ctx, commandNames).
					Return(commandNames[1], nil)

				return &m{cliScript}
			},
		},
		{
			"fail with unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
				commands: map[string]func() (string, error){
					commandNames[0]: func() (string, error) { return "key1", nil },
					commandNames[1]: func() (string, error) { return "key2", nil },
				},
			},
			testError,
			func(a *args) *m {

				cliScript := mocks.NewCLISelectCommandSupporter(t)
				cliScript.EXPECT().
					SelectCommand(a.ctx, commandNames).
					Return("", testError)

				return &m{cliScript}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			result, err := SelectCommand(
				tt.args.ctx,
				m.cliScript,
				tt.args.commands,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			r1, _ := tt.args.commands[commandNames[1]]()
			r2, _ := result()
			assert.Equal(t, r1, r2)
		})
	}
}
