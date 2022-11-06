package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type App struct {
	Config            *Config
	Logger            port.Logger
	Server            port.Server
	Encryptor         port.EncryptDecryptor
	Hasher            port.PasswordHasher
	Storage           port.Storage
	Transactor        port.Transactor
	UserRepository    port.UserRepository
	TypeRepository    port.TypeRepository
	EntryRepository   port.EntryRepository
	SessionRepository port.SessionRepository
}

func NewApp(
	config *Config,
	logger port.Logger,
	server port.Server,
	encryptor port.EncryptDecryptor,
	hasher port.PasswordHasher,
	storage port.Storage,
	transactor port.Transactor,
	userRepository port.UserRepository,
	typeRepository port.TypeRepository,
	entryRepository port.EntryRepository,
	sessionRepository port.SessionRepository,
) *App {
	return &App{
		Config:            config,
		Logger:            logger,
		Server:            server,
		Encryptor:         encryptor,
		Hasher:            hasher,
		Storage:           storage,
		Transactor:        transactor,
		UserRepository:    userRepository,
		TypeRepository:    typeRepository,
		EntryRepository:   entryRepository,
		SessionRepository: sessionRepository,
	}
}

func (a *App) Run(handlerCollection any) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	shutdownCh := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := a.Server.Shutdown(); err != nil {
			a.Logger.Error(fmt.Errorf("GRPC server Shutdown: %w", err))
		}

		close(shutdownCh)
	}()

	a.Logger.Info("starting GRPC server", "addr", a.Config.RunAddress)

	err := a.Server.Serve(handlerCollection)
	if err != nil {
		return err
	}

	<-shutdownCh
	fmt.Println("GRPC server Shutdown gracefully")

	return nil
}
