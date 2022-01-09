package Config

import (
	"go.uber.org/zap/zapcore"
	"log"
)

type ConfigLoader interface {
	load() Cfg
}

type Cfg struct {
	Service struct {
		WorkersCount int           `yaml:"WorkersCount"`
		LogLevel     zapcore.Level `yaml:"LogLevel"`
	}
}

func NewConfig() *Cfg {
	log.Println("Load yaml config file...")
	cfg := yamlConfig{}.load() // TODO check read yaml
	log.Println("WorkersCount=", cfg.Service.WorkersCount)
	return &cfg
}
