package port

import (
	"context"
	"database/sql"
)

type ConfigLoader interface {
	Load(cfg interface{}) error
	LoadFile(path string, cfg interface{}) error
}

type Logger interface {
	EnableDebug() error
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(err error, args ...interface{})
	Fatal(err error, args ...interface{})
}

type Server interface {
	Serve(handlerCollection any) error
	Shutdown() error
}

type PasswordHasher interface {
	Hash(password string) string
}

type Encryptor interface {
	Encrypt(data []byte) (string, error)
}

type Decryptor interface {
	Decrypt(encrypted string) ([]byte, error)
}

type EncryptDecryptor interface {
	Encryptor
	Decryptor
}

type Storage interface {
	Connect() *sql.DB
	Initialization(migrationDir string) error
	Ping(ctx context.Context) error
}

type TransactSupporter interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
}

type RowQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Querier interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Transactor interface {
	TransactSupporter
	Querier
	RowQuerier
	Execer
}
