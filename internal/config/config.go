package config

import (
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

var config *Cfg

type ConfigLoader interface {
	load() (Cfg, error)
}

type Cfg struct {
	Service struct {
		WorkersCount   int           `yaml:"WorkersCount" envconfig:"WORKERS_COUNT"`
		LogLevel       zapcore.Level `yaml:"LogLevel" envconfig:"LOG_LEVEL"`
		LinksLivesDays int           `yaml:"LinksLivesDays" envconfig:"LINKS_LIVES_DAYS"`
	}
	Server struct {
		Http struct {
			Port    string        `yaml:"port" envconfig:"PORT"`
			Timeout time.Duration `yaml:"timeout" envconfig:"TIMEOUT"`
		}
	}
	Data struct {
		Database struct {
			Driver string `yaml:"driver" envconfig:"DB_DRIVER"`
			Source string `yaml:"source" envconfig:"DATABASE_URL"`
		}
	}
}

func NewConfig() *Cfg {
	log.Println("Load yaml config file...")
	cfg, err := yamlConfig{}.load()
	if err != nil {
		log.Println("Error load config from yaml. Try from ENV")
		cfg, err = envConfig{}.load()
		if err != nil {
			log.Panicf("%v", err)
		}
	}

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
