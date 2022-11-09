package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTypes(t *testing.T) {
	type m struct {
		typeRepository *mocks.TypeAllGetter
	}

	testTypes := []*entity.Type{
		{
			ID:       uuid.New(),
			Name:     "type1",
			IsBinary: false,
		},
		{
			ID:       uuid.New(),
			Name:     "type2",
			IsBinary: true,
		},
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
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID: uuid.New(),
				}),
			},
			nil,
			func(a *args) *m {
				typeRepository := mocks.NewTypeAllGetter(t)
				typeRepository.EXPECT().
					GetAll(a.ctx).
					Return(testTypes, nil)

				return &m{typeRepository}
			},
		},
		{
			"fail with type not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID: uuid.New(),
				}),
			},
			ErrTypeNotFound,
			func(a *args) *m {
				typeRepository := mocks.NewTypeAllGetter(t)
				typeRepository.EXPECT().
					GetAll(a.ctx).
					Return(nil, ErrTypeNotFound)

				return &m{typeRepository}
			},
		},
		{
			"fail with unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID: uuid.New(),
				}),
			},
			ErrInternalServerError,
			func(a *args) *m {
				typeRepository := mocks.NewTypeAllGetter(t)
				typeRepository.EXPECT().
					GetAll(a.ctx).
					Return(nil, fmt.Errorf("unexpected error"))

				return &m{typeRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			types, err := GetAllTypes(
				tt.args.ctx,
				m.typeRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testTypes[0].Name, types[0].Name)
		})
	}
}
