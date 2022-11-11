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

func TestList(t *testing.T) {
	type m struct {
		client    *mocks.GRPCClientListSupporter
		cliScript *mocks.CLIListOfEntriesSupporter
	}

	testError := fmt.Errorf("unexpected error")
	testTypes := []*model.Type{
		{ID: "id1", Name: "type1"},
		{ID: "id2", Name: "type2"},
		{ID: "id3", Name: "type3"},
	}
	testEntries := []*model.Entry{
		{ID: "id1", Name: "entry1", TypeID: "id1"},
		{ID: "id2", Name: "entry2", TypeID: "id2"},
		{ID: "id3", Name: "entry3", TypeID: "id3"},
	}
	testEntryIndex := 1

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
				client := mocks.NewGRPCClientListSupporter(t)
				client.EXPECT().
					GetAllEntries(a.ctx).
					Return(testEntries, nil)
				client.EXPECT().
					GetEntry(a.ctx, testEntries[testEntryIndex].ID).
					Return(testEntries[testEntryIndex], nil)

				cliScript := mocks.NewCLIListOfEntriesSupporter(t)
				cliScript.EXPECT().
					ListOfEntries(a.ctx, []string{testEntries[0].Name, testEntries[1].Name, testEntries[2].Name}).
					Return(testEntryIndex, nil)

				return &m{client, cliScript}
			},
		},
		{
			"fail with GetAllEntries unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {
				client := mocks.NewGRPCClientListSupporter(t)
				client.EXPECT().
					GetAllEntries(a.ctx).
					Return(nil, testError)

				cliScript := mocks.NewCLIListOfEntriesSupporter(t)

				return &m{client, cliScript}
			},
		},
		{
			"fail with ListOfEntries unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {
				client := mocks.NewGRPCClientListSupporter(t)
				client.EXPECT().
					GetAllEntries(a.ctx).
					Return(testEntries, nil)

				cliScript := mocks.NewCLIListOfEntriesSupporter(t)
				cliScript.EXPECT().
					ListOfEntries(a.ctx, []string{testEntries[0].Name, testEntries[1].Name, testEntries[2].Name}).
					Return(0, testError)

				return &m{client, cliScript}
			},
		},
		{
			"fail with GetEntry unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, "session key"),
			},
			testError,
			func(a *args) *m {
				client := mocks.NewGRPCClientListSupporter(t)
				client.EXPECT().
					GetAllEntries(a.ctx).
					Return(testEntries, nil)
				client.EXPECT().
					GetEntry(a.ctx, testEntries[testEntryIndex].ID).
					Return(nil, testError)

				cliScript := mocks.NewCLIListOfEntriesSupporter(t)
				cliScript.EXPECT().
					ListOfEntries(a.ctx, []string{testEntries[0].Name, testEntries[1].Name, testEntries[2].Name}).
					Return(testEntryIndex, nil)

				return &m{client, cliScript}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			result, err := List(
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
			assert.Equal(t, fmt.Sprintf(
				"ID:\t%s\nName:\t%s\nType:\t%s\nMetadata:\n%s\nData:\n%s",
				testEntries[testEntryIndex].ID,
				testEntries[testEntryIndex].Name,
				testTypes[testEntryIndex].Name,
				"",
				"",
			), result)
		})
	}
}
