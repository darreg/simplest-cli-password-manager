package handler

import (
	"github.com/alrund/yp-2-project/server/internal/application/app"
	"github.com/alrund/yp-2-project/server/pkg/proto"
)

// Collection of GRPC handlers.
type Collection struct {
	proto.UnimplementedAppServer
	a *app.App
}

func NewCollection(a *app.App) *Collection {
	return &Collection{a: a}
}
