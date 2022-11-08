package port

import (
	"context"

	"github.com/alrund/yp-2-project/server/internal/domain/entity"
	"github.com/google/uuid"
)

type SessionGetter interface {
	Get(ctx context.Context, sessionID uuid.UUID) (*entity.Session, error)
}

type SessionAdder interface {
	Add(ctx context.Context, session *entity.Session) error
}

type SessionChanger interface {
	Change(ctx context.Context, session *entity.Session) error
}

type SessionRemover interface {
	Remove(ctx context.Context, sessionID uuid.UUID) error
}

type SessionRefresher interface {
	SessionGetter
	SessionChanger
}

type SessionRepository interface {
	SessionGetter
	SessionAdder
	SessionChanger
	SessionRemover
}
