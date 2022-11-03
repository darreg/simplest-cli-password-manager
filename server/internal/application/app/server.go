package app

import (
	"net/http"
)

type SessionContextKey string

func (a *App) Serve() error {
	a.Logger.Info("starting GRPC server", "addr", a.Config.RunAddress)
	return http.ListenAndServe(a.Config.RunAddress, a.Router)
}
