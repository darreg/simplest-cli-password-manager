package builder

import (
	"github.com/alrund/yp-2-project/server/internal/application/app"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/adapter"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/repository"
)

// Builder builds the application structure.
func Builder(config *app.Config, logger port.Logger) (*app.App, error) {
	if config.Debug {
		err := logger.EnableDebug()
		if err != nil {
			return nil, err
		}
	}

	storage, err := adapter.NewStorage(config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	err = storage.Initialization(config.MigrationDir)
	if err != nil {
		return nil, err
	}

	var (
		hasher            = adapter.NewHasher()
		encryptor         = adapter.NewEncryptor(config.CipherPass)
		transactor        = adapter.NewTransactor(storage.Connect())
		userRepository    = repository.NewUserRepository(transactor)
		typeRepository    = repository.NewTypeRepository(transactor)
		entryRepository   = repository.NewEntryRepository(transactor)
		sessionRepository = repository.NewSessionRepository(transactor)
		server            = adapter.NewServer(
			config.SessionLifeTime,
			config.RunAddress,
			config.CertFile,
			config.KeyFile,
			encryptor,
			sessionRepository,
			logger,
		)
	)

	return app.NewApp(
		config,
		logger,
		server,
		encryptor,
		hasher,
		storage,
		transactor,
		userRepository,
		typeRepository,
		entryRepository,
		sessionRepository,
	), nil
}
