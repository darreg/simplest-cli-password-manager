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

func TestSet(t *testing.T) {
	type m struct {
		client    *mocks.GRPCClientSetSupporter
		cliScript *mocks.CLISetEntrySupporter
	}

	testError := fmt.Errorf("unexpected error")
	testTypes := []*model.Type{
		{ID: "id1", Name: "type1"},
		{ID: "id2", Name: "type2"},
		{ID: "id3", Name: "type3"},
	}
	testTypeNames := []string{testTypes[0].Name, testTypes[1].Name, testTypes[2].Name}
	testDto := &SetEntryDTO{
		TypeIndex: 2,
		Name:      "Entry Name",
		Metadata:  "Entry Metadata",
		Data:      "Entry Data",
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
				cliScript := mocks.NewCLISetEntrySupporter(t)
				cliScript.EXPECT().
					SetEntry(a.ctx, testTypeNames).
					Return(testDto, nil)

				client := mocks.NewGRPCClientSetSupporter(t)
				client.EXPECT().
					SetEntry(
						a.ctx,
						testTypes[testDto.TypeIndex].ID,
						testDto.Name,
						testDto.Metadata,
						[]byte(testDto.Data),
					).
					Return(nil)

				return &m{client, cliScript}
			},
		},
		{
			"fail with cliScript.SetEntry unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {
				cliScript := mocks.NewCLISetEntrySupporter(t)
				cliScript.EXPECT().
					SetEntry(a.ctx, testTypeNames).
					Return(nil, testError)

				client := mocks.NewGRPCClientSetSupporter(t)

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
				cliScript := mocks.NewCLISetEntrySupporter(t)
				cliScript.EXPECT().
					SetEntry(a.ctx, testTypeNames).
					Return(&struct{}{}, nil)

				client := mocks.NewGRPCClientSetSupporter(t)

				return &m{client, cliScript}
			},
		},
		{
			"fail with client.SetEntry unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {
				cliScript := mocks.NewCLISetEntrySupporter(t)
				cliScript.EXPECT().
					SetEntry(a.ctx, testTypeNames).
					Return(testDto, nil)

				client := mocks.NewGRPCClientSetSupporter(t)
				client.EXPECT().
					SetEntry(
						a.ctx,
						testTypes[testDto.TypeIndex].ID,
						testDto.Name,
						testDto.Metadata,
						[]byte(testDto.Data),
					).
					Return(testError)

				return &m{client, cliScript}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			_, err := Set(
				tt.args.ctx,
				m.client,
				m.cliScript,
				testTypes,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
