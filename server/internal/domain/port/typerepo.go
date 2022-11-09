package port

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/google/uuid"
)

type TypeGetter interface {
	Get(ctx context.Context, tpID uuid.UUID) (*entity.Type, error)
}

type TypeAllGetter interface {
	GetAll(ctx context.Context) ([]*entity.Type, error)
}

type TypeAdder interface {
	Add(ctx context.Context, tp *entity.Type) error
}

type TypeChanger interface {
	Change(ctx context.Context, tp *entity.Type) error
}

type TypeRemover interface {
	Remove(ctx context.Context, tpID uuid.UUID) error
}

type TypeRepository interface {
	TypeGetter
	TypeAllGetter
	TypeAdder
	TypeChanger
	TypeRemover
}
