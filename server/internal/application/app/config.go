package app

import (
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type Config struct {
	Debug        bool   `env-default:"false"`
	MigrationDir string `env-default:"migrations"`
	RunAddress   string `env:"RUN_ADDRESS" env-default:"localhost:3000"`
	DatabaseURI  string `env:"DATABASE_URI" env-default:"postgres://dev:dev@localhost:5432/dev?sslmode=disable"`
	CipherPass   string `env:"CIPHER_PASSWORD" env-default:"J53RPX6"`
	CertFile     string `env:"CERT_FILE" json:"cert_file"`
	KeyFile      string `env:"KEY_FILE" json:"key_file"`
}

func NewConfig(loader port.ConfigLoader) (*Config, error) {
	cfg := &Config{}

	flags := NewFlags()
	cfg.Debug = flags.Debug
	cfg.RunAddress = flags.A
	cfg.DatabaseURI = flags.D

	if err := loader.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
