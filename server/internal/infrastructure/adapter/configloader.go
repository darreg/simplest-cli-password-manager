package adapter

import "github.com/ilyakaznacheev/cleanenv"

type ConfigLoader struct{}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (c ConfigLoader) Load(cfg interface{}) error {
	return cleanenv.ReadEnv(cfg)
}

func (c ConfigLoader) LoadFile(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}
