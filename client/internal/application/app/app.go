package app

import (
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type App struct {
	Config *Config
	Logger port.Logger
}

func NewApp(
	config *Config,
	logger port.Logger,
) *App {
	return &App{
		Config: config,
		Logger: logger,
	}
}

func (a *App) Run() error {
	return nil
}
