//go:build unit

package usecase

import (
	"context"
	"fmt"
	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSetEntry(t *testing.T) {
	type m struct {
		encryptor       *mocks.Encryptor
		entryRepository *mocks.EntryAdder
		userRepository  *mocks.UserGetter
		typeRepository  *mocks.TypeGetter
	}

	testUserID := uuid.New()
	testTypeID := uuid.New()
	testData := []byte("test data")

	type args struct {
		ctx      context.Context
		entryDTO *SetEntryDTO
		userID   uuid.UUID
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
				entryDTO: &SetEntryDTO{
					TypeID:   testTypeID.String(),
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
				userID: testUserID,
			},
			nil,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				typeRepository := mocks.NewTypeGetter(t)
				typeRepository.EXPECT().
					Get(a.ctx, testTypeID).
					Return(&entity.Type{
						ID: testTypeID,
					}, nil)

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return("", nil)

				entryRepository := mocks.NewEntryAdder(t)
				entryRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Entry")).
					Return(nil)

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
		{
			"fail with incorrect session",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, nil),
				entryDTO: &SetEntryDTO{
					TypeID:   testTypeID.String(),
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
			},
			ErrIncorrectSession,
			func(a *args) *m {
				typeRepository := mocks.NewTypeGetter(t)
				userRepository := mocks.NewUserGetter(t)
				encryptor := mocks.NewEncryptor(t)
				entryRepository := mocks.NewEntryAdder(t)

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
		{
			"fail with entry uuid parse error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				entryDTO: &SetEntryDTO{
					TypeID:   "incorrect",
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
			},
			ErrInvalidArgument,
			func(a *args) *m {
				typeRepository := mocks.NewTypeGetter(t)
				userRepository := mocks.NewUserGetter(t)
				encryptor := mocks.NewEncryptor(t)
				entryRepository := mocks.NewEntryAdder(t)

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
		{
			"fail with type not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				entryDTO: &SetEntryDTO{
					TypeID:   testTypeID.String(),
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
			},
			ErrTypeNotFound,
			func(a *args) *m {
				typeRepository := mocks.NewTypeGetter(t)
				typeRepository.EXPECT().
					Get(a.ctx, testTypeID).
					Return(nil, ErrTypeNotFound)
				userRepository := mocks.NewUserGetter(t)
				encryptor := mocks.NewEncryptor(t)
				entryRepository := mocks.NewEntryAdder(t)

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
		{
			"fail with user not found",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				entryDTO: &SetEntryDTO{
					TypeID:   testTypeID.String(),
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
				userID: testUserID,
			},
			ErrUserNotFound,
			func(a *args) *m {
				typeRepository := mocks.NewTypeGetter(t)
				typeRepository.EXPECT().
					Get(a.ctx, testTypeID).
					Return(&entity.Type{
						ID: testTypeID,
					}, nil)

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(nil, ErrUserNotFound)
				encryptor := mocks.NewEncryptor(t)
				entryRepository := mocks.NewEntryAdder(t)

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
		{
			"fail with encrypt error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				entryDTO: &SetEntryDTO{
					TypeID:   testTypeID.String(),
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
				userID: testUserID,
			},
			ErrInternalServerError,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				typeRepository := mocks.NewTypeGetter(t)
				typeRepository.EXPECT().
					Get(a.ctx, testTypeID).
					Return(&entity.Type{
						ID: testTypeID,
					}, nil)

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return("", fmt.Errorf("unexpected error"))

				entryRepository := mocks.NewEntryAdder(t)

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
		{
			"fail with entry repository unexpected error",
			&args{
				ctx: context.WithValue(context.Background(), port.SessionContextKey, &entity.Session{
					ID:     uuid.New(),
					UserID: testUserID,
				}),
				entryDTO: &SetEntryDTO{
					TypeID:   testTypeID.String(),
					Name:     "",
					Metadata: "",
					Data:     testData,
				},
				userID: testUserID,
			},
			ErrInternalServerError,
			func(a *args) *m {
				testUser := &entity.User{
					ID: a.userID,
				}

				typeRepository := mocks.NewTypeGetter(t)
				typeRepository.EXPECT().
					Get(a.ctx, testTypeID).
					Return(&entity.Type{
						ID: testTypeID,
					}, nil)

				userRepository := mocks.NewUserGetter(t)
				userRepository.EXPECT().
					Get(a.ctx, a.userID).
					Return(testUser, nil)

				encryptor := mocks.NewEncryptor(t)
				encryptor.EXPECT().
					Encrypt(mock.AnythingOfType("[]uint8")).
					Return("", nil)

				entryRepository := mocks.NewEntryAdder(t)
				entryRepository.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Entry")).
					Return(fmt.Errorf("unexpected error"))

				return &m{encryptor, entryRepository, userRepository, typeRepository}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			err := SetEntry(
				tt.args.ctx,
				tt.args.entryDTO,
				m.encryptor,
				m.entryRepository,
				m.userRepository,
				m.typeRepository,
			)

			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
