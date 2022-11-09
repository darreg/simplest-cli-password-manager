package app

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func (a *App) Run(client port.GRPCClientSupporter, cliScript port.CLIScriptSupporter) error {
	cred, err := credentials.NewClientTLSFromFile(a.Config.CertFile, "")
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(a.Config.ServerAddress, grpc.WithTransportCredentials(cred))
	if err != nil {
		return err
	}
	defer conn.Close()

	err = client.SetGRPCClient(proto.NewAppClient(conn))
	if err != nil {
		return err
	}

	for {
		if client.IsEmptySessionKey() {
			err := usecase.Login(
				context.Background(),
				client,
				cliScript,
			)
			if err != nil {
				return err
			}
		}

		err = usecase.Command(
			context.Background(),
			client,
			cliScript,
		)
		if err != nil {
			return err
		}
	}
}
