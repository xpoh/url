package config

import (
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

type ConfigLoader interface {
	load() Cfg
}

type Cfg struct {
	Service struct {
		WorkersCount   int           `yaml:"WorkersCount"`
		LogLevel       zapcore.Level `yaml:"LogLevel"`
		LinksLivesDays int           `yaml:"LinksLivesDays"`
	}
	Server struct {
		Http struct {
			Addr    string        `yaml:"addr"`
			Timeout time.Duration `yaml:"timeout"`
		}
	}
}

func NewConfig() *Cfg {
	log.Println("Load yaml config file...")
	cfg := yamlConfig{}.load() // TODO check read yaml
	log.Println("WorkersCount=", cfg.Service.WorkersCount)
	log.Printf("%v\n", cfg)
	return &cfg
}
