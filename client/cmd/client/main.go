package main

import (
	"fmt"
	"github.com/alrund/yp-2-project/client/internal/application/app"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/adapter"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/builder"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/cli"
	"github.com/alrund/yp-2-project/client/internal/infrastructure/client"
	"golang.org/x/net/context"
)

const defaultBuildValue string = "N/A"

var (
	buildVersion = defaultBuildValue
	buildDate    = defaultBuildValue
)

func main() {
	printBuildInfo()

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

func printBuildInfo() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
}
