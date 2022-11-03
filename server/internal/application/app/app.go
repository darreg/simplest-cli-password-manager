package app

import "github.com/alrund/yp-2-project/server/internal/domain/port"

type App struct {
	Config          *Config
	Logger          port.Logger
	Encryptor       port.Encryptor
	Hasher          port.PasswordHasher
	Storage         port.Storage
	Transactor      port.Transactor
	UserRepository  port.UserRepository
	TypeRepository  port.TypeRepository
	EntryRepository port.EntryRepository
}

func NewApp(
	config *Config,
	logger port.Logger,
	encryptor port.Encryptor,
	hasher port.PasswordHasher,
	storage port.Storage,
	transactor port.Transactor,
	userRepository port.UserRepository,
	typeRepository port.TypeRepository,
	entryRepository port.EntryRepository,
) *App {
	return &App{
		Config:          config,
		Logger:          logger,
		Encryptor:       encryptor,
		Hasher:          hasher,
		Storage:         storage,
		Transactor:      transactor,
		UserRepository:  userRepository,
		TypeRepository:  typeRepository,
		EntryRepository: entryRepository,
	}
}

func (a *App) Run() error {
	return a.Serve()
}

func (a *App) Stop() error {
	return nil
}
