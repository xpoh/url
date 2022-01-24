package config

import (
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var config *Cfg

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
	Data struct {
		Database struct {
			Driver string `yaml:"driver"`
			Source string `yaml:"source"`
		}
	}
}

func NewConfig() *Cfg {
	log.Println("Load yaml config file...")
	cfg := yamlConfig{}.load()
	log.Println("WorkersCount=", cfg.Service.WorkersCount)
	log.Printf("%v\n", cfg)
	config = &cfg
	return config
}

func GetConfig() *Cfg {
	if config == nil {
		panic("get nil config error!!!")
	}
	return config
}
