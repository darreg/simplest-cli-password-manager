package main

import (
	"github.com/alrund/yp-2-project/client/internal/application/app"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/adapter"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/builder"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/cli"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/client"
	"golang.org/x/net/context"
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

	if err := a.Run(context.Background(), client.New(), cli.New()); err != nil {
		logger.Fatal(err)
	}
}
