package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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

// Run starts the application.
func (a *App) Run(ctx context.Context, client port.GRPCClientSupporter, cliScript port.CLIScriptSupporter) error {
	cred, err := credentials.NewClientTLSFromFile(a.Config.CertFile, "")
	if err != nil {
		return err
	}

	conn, err := grpc.DialContext(ctx, a.Config.ServerAddress, grpc.WithTransportCredentials(cred))
	if err != nil {
		return err
	}
	defer conn.Close()

	err = client.SetGRPCClient(proto.NewAppClient(conn))
	if err != nil {
		return err
	}

	types, err := client.GetAllTypes(ctx)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		fmt.Println("\nThe application is stopped by the user.\nBye!")
	}()

	err = a.Wait(ctx, client, cliScript, types)
	if err != nil {
		return err
	}

	return nil
}

// Wait waiting for the command to be entered.
func (a *App) Wait(
	ctx context.Context,
	client port.GRPCClientSupporter,
	cliScript port.CLIScriptSupporter,
	types []*model.Type,
) error {
	var err error
	for {
		if client.IsEmptySessionKey() {
			err = a.Login(ctx, client, cliScript, map[string]func() (string, error){
				"Login":        func() (string, error) { return usecase.Login(ctx, client, cliScript) },
				"Registration": func() (string, error) { return usecase.Registration(ctx, client, cliScript) },
			})
			if err != nil {
				return err
			}
			continue
		}

		err = a.Command(ctx, cliScript, map[string]func() (string, error){
			"List": func() (string, error) { return usecase.List(ctx, client, cliScript, types) },
			"Set":  func() (string, error) { return usecase.Set(ctx, client, cliScript, types) },
			"Exit": func() (string, error) { return usecase.Exit(ctx) },
		})
		if err != nil {
			return err
		}
	}
}

// Login authorizes and outputs a greeting.
func (a *App) Login(
	ctx context.Context,
	client port.GRPCClientSupporter,
	cliScript port.CLIScriptSupporter,
	loginMethodFns map[string]func() (string, error),
) error {
	loginMethodFn, err := usecase.SelectLoginMethod(ctx, cliScript, loginMethodFns)
	if err != nil {
		return err
	}

	sessionKey, err := loginMethodFn()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.Unauthenticated {
				fmt.Printf("Check your credentials\n\n")
				return nil
			}
		} else if errors.Is(err, usecase.ErrIncorrectPassword) {
			fmt.Printf("X Passwords must be the same\n\n")
			return nil
		}
		return err
	}

	err = client.SetSessionKey(sessionKey)
	if err != nil {
		return err
	}

	result, err := usecase.Greetings(ctx, client)
	if err != nil {
		return err
	}

	if result != "" {
		fmt.Println(result)
	}

	return nil
}

// Command processes user commands.
func (a *App) Command(
	ctx context.Context,
	cliScript port.CLIScriptSupporter,
	commandFns map[string]func() (string, error),
) error {
	commandFn, err := usecase.SelectCommand(ctx, cliScript, commandFns)
	if err != nil {
		return err
	}

	result, err := commandFn()
	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.NotFound {
				fmt.Printf("No entries\n\n")
				return nil
			} else if e.Code() == codes.Unauthenticated {
				fmt.Printf("Check your credentials\n\n")
				return nil
			}
		}
		return err
	}

	if result != "" {
		fmt.Println(result)
	}

	return nil
}
