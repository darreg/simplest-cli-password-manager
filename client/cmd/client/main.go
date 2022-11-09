package main

import (
	"github.com/alrund/yp-2-project/client/internal/application/app"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/adapter"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/builder"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/client"
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

	if err := a.Run(client.New(a)); err != nil {
		logger.Fatal(err)
	}
}
