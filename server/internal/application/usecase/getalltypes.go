package usecase

import (
	"context"
	"errors"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

// GetAllTypes gets all types.
func GetAllTypes(
	ctx context.Context,
	typeRepository port.TypeAllGetter,
) ([]*entity.Type, error) {
	types, err := typeRepository.GetAll(ctx)
	if err != nil {
		if errors.Is(err, ErrTypeNotFound) {
			return nil, ErrTypeNotFound
		}
		return nil, ErrInternalServerError
	}

	return types, nil
}
