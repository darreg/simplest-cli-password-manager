package port

import (
	"context"
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

type Client interface {
	SetGRPCClient(client any) error
	Registration(ctx context.Context) error
	Login(ctx context.Context) error
}
