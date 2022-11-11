package usecase

import (
	"context"
	"time"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/google/uuid"
)

type SetEntryDTO struct {
	TypeID   string
	Name     string
	Metadata string
	Data     []byte
}

// SetEntry adds an entry.
func SetEntry(
	ctx context.Context,
	entryDTO *SetEntryDTO,
	encryptor port.Encryptor,
	entryRepository port.EntryAdder,
	userRepository port.UserGetter,
	typeRepository port.TypeGetter,
) error {
	contextSession := ctx.Value(port.SessionContextKey)
	session, ok := contextSession.(*entity.Session)
	if !ok {
		return ErrIncorrectSession
	}

	typeID, err := uuid.Parse(entryDTO.TypeID)
	if err != nil {
		return ErrInvalidArgument
	}

	tp, err := typeRepository.Get(ctx, typeID)
	if err != nil {
		return ErrTypeNotFound
	}

	user, err := userRepository.Get(ctx, session.UserID)
	if err != nil {
		return ErrUserNotFound
	}

	encryptedData, err := encryptor.Encrypt(entryDTO.Data)
	if err != nil {
		return ErrInternalServerError
	}

	nowTime := time.Now()
	entry := entity.NewEntry(
		user.ID,
		tp.ID,
		entryDTO.Name,
		entryDTO.Metadata,
		[]byte(encryptedData),
		&nowTime,
		&nowTime,
	)
	err = entryRepository.Add(ctx, entry)
	if err != nil {
		return ErrInternalServerError
	}

	return nil
}
