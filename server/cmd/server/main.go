package main

import (
	"github.com/alrund/yp-2-project/server/internal/application/app"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/builder"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/handler"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger := adapter.NewLogger()

	config, err := app.NewConfig(adapter.NewConfigLoader())
	if err != nil {
		logger.Fatal(err)
	}

	a, err := builder.Builder(config, logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err := a.Run(handler.NewCollection(a)); err != nil {
		logger.Fatal(err)
	}
}
