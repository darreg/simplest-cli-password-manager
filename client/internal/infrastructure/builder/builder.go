package builder

import (
	"github.com/alrund/yp-2-project/client/internal/application/app"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

// Builder builds the application structure.
func Builder(config *app.Config, logger port.Logger) (*app.App, error) {
	if config.Debug {
		err := logger.EnableDebug()
		if err != nil {
			return nil, err
		}
	}

	return app.NewApp(
		config,
		logger,
	), nil
}
