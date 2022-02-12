package config

import (
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct{}

func (e envConfig) load() (Cfg, error) {
	cfg := Cfg{}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
