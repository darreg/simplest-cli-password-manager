//go:build unit

package usecase

import (
	"context"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	type m struct {
		userRepository *mocks.UserGetter
	}

	testUserID := uuid.New()

	type args struct {
		ctx    context.Context
		userID uuid.UUID
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
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				userID: testUserID,
			},
			nil,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				return &m{userGetter}
			},
		},
		{
			"fail with not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				userID: testUserID,
			},
			ErrUserNotFound,
			func(a *args) *m {

				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(nil, ErrUserNotFound)

				return &m{userGetter}
			},
		},
		{
			"fail with incorrect session",
			&args{
				ctx:    context.WithValue(context.Background(), port.SessionContextKey, nil),
				userID: testUserID,
			},
			ErrIncorrectSession,
			func(a *args) *m {

				userGetter := mocks.NewUserGetter(t)

				return &m{userGetter}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			user, err := GetUser(
				tt.args.ctx,
				m.userRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.args.userID.String(), user.ID.String())
		})
	}
}
