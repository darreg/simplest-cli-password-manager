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

type GRPCClientRegistrationSupporter interface {
	Registration(ctx context.Context, name, login, password string) (string, error)
}

type GRPCClientLoginSupporter interface {
	Login(ctx context.Context, login, password string) (string, error)
}

type GRPCClientListSupporter interface {
	GetEntry(ctx context.Context, entryID string) (*model.Entry, error)
	GetAllEntries(ctx context.Context) ([]*model.Entry, error)
}

type GRPCClientSetSupporter interface {
	SetEntry(ctx context.Context, typeID, name, metadata string, data []byte) error
}

type GRPCClientSupporter interface {
	SetGRPCClient(client any) error
	SetSessionKey(sessionKey string) error
	IsEmptySessionKey() bool
	GetAllTypes(ctx context.Context) ([]*model.Type, error)

	GRPCClientRegistrationSupporter
	GRPCClientLoginSupporter
	GRPCClientListSupporter
	GRPCClientSetSupporter

	// RemoveEntry(ctx context.Context, entryID string) error
}

type CLISelectLoginMethodSupporter interface {
	SelectLoginMethod(ctx context.Context, options []string, data any) error
}

type CLILoginSupporter interface {
	Login(ctx context.Context, data any) error
}

type CLIRegistrationSupporter interface {
	Registration(ctx context.Context, data any) error
}

type CLISelectCommandSupporter interface {
	SelectCommand(ctx context.Context, options []string, data any) error
}

type CLISetEntrySupporter interface {
	SetEntry(ctx context.Context, types []string, data any) error
}

type CLIListOfEntriesSupporter interface {
	ListOfEntries(ctx context.Context, entries []string, data any) error
}

type CLIScriptSupporter interface {
	CLISelectLoginMethodSupporter
	CLILoginSupporter
	CLIRegistrationSupporter
	CLISelectCommandSupporter
	CLISetEntrySupporter
	CLIListOfEntriesSupporter
}
