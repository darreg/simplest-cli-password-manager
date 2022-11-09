package app

import (
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type Config struct {
	Debug         bool   `env-default:"false"`
	ServerAddress string `env:"SERVER_ADDRESS" env-default:"localhost:3000"`
	CertFile      string `env:"CERT_FILE" env-default:"local.crt"`
}

func NewConfig(loader port.ConfigLoader) (*Config, error) {
	cfg := &Config{}

	flags := NewFlags()
	cfg.Debug = flags.Debug

	if err := loader.Load(cfg); err != nil {
		return nil, err
	}

	ReadFlags(flags, cfg)

	return cfg, nil
}

func ReadFlags(f *Flags, cfg *Config) {
	if f.A != NotAvailable {
		cfg.ServerAddress = f.A
	}
}
