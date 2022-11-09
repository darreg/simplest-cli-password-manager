package port

import (
	"context"
	"github.com/alrund/yp-2-project/client/internal/domain/model"
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

type TypeDecoder interface {
	Decode(data []byte) []byte
}

type GRPCClientLoginMethodSupporter interface {
	Registration(ctx context.Context, login, password string) (string, error)
	Login(ctx context.Context, login, password string) (string, error)
	SetSessionKey(sessionKey string) error
}

type GRPCClientSupporter interface {
	SetGRPCClient(client any) error
	IsEmptySessionKey() bool

	GRPCClientLoginMethodSupporter

	SetEntry(ctx context.Context, typeID, name, metadata string, data []byte) error
	GetEntry(ctx context.Context, entryID string) (*model.Entry, error)
	GetAllEntries(ctx context.Context) ([]*model.Entry, error)
	GetAllTypes(ctx context.Context) ([]*model.Type, error)
	// RemoveEntry(ctx context.Context, entryID string) error
}

type CLILoginMethodSupporter interface {
	SelectLoginMethod(ctx context.Context, options []string, data any) error
	Login(ctx context.Context, data any) error
	Registration(ctx context.Context, data any) error
}

type CLICommandSupporter interface {
	SelectCommand(ctx context.Context, options []string, data any) error
	SetEntry(ctx context.Context, types []string, data any) error
	ListOfEntries(ctx context.Context, entries []string, data any) error
}

type CLIScriptSupporter interface {
	CLILoginMethodSupporter
	CLICommandSupporter
}
