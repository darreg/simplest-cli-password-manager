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

func (a *App) Run(ctx context.Context, client port.GRPCClientSupporter, cliScript port.CLIScriptSupporter) error {
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
		err = Login(ctx, client, cliScript)
		if err != nil {
			return err
		}

		err = Command(ctx, client, cliScript)
		if err != nil {
			return err
		}
	}
}

func Login(ctx context.Context, client port.GRPCClientSupporter, cliScript port.CLIScriptSupporter) error {
	const (
		Login        string = "Login"
		Registration string = "Registration"
	)

	var (
		sessionKey   string
		loginMethods = []string{Login, Registration}
	)

	if client.IsEmptySessionKey() {
		loginMethodIndex, err := usecase.SelectLoginMethod(ctx, cliScript, loginMethods)
		if err != nil {
			return err
		}

		switch loginMethods[loginMethodIndex] {
		case Login:
			sessionKey, err = usecase.Login(ctx, client, cliScript)
		case Registration:
			sessionKey, err = usecase.Registration(ctx, client, cliScript)
		}
		if err != nil {
			return err
		}

		err = client.SetSessionKey(sessionKey)
		if err != nil {
			return err
		}
	}

	return nil
}

func Command(ctx context.Context, client port.GRPCClientSupporter, cliScript port.CLIScriptSupporter) error {
	const (
		List string = "List"
		Set  string = "Set"
	)

	commands := []string{List, Set}

	types, err := client.GetAllTypes(ctx)
	if err != nil {
		return err
	}

	commandIndex, err := usecase.SelectCommand(ctx, cliScript, commands)
	if err != nil {
		return err
	}

	switch commands[commandIndex] {
	case List:
		err = usecase.List(ctx, client, cliScript, types)
	case Set:
		err = usecase.Set(ctx, client, cliScript, types)
	}
	if err != nil {
		return err
	}

	return nil
}
