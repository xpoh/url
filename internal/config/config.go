// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config implement extern configuration
// for butty service
package config

import (
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

// config global variable
var config *Cfg

// ConfigLoader loader config interface
type ConfigLoader interface {
	// load config from any source
	// return Cfg structure or error
	load() (Cfg, error)
}

// Cfg is a main config structure
type Cfg struct {
	Service struct {
		WorkersCount   int           `yaml:"WorkersCount" envconfig:"WORKERS_COUNT"`
		LogLevel       zapcore.Level `yaml:"LogLevel" envconfig:"LOG_LEVEL"`
		LinksLivesDays int           `yaml:"LinksLivesDays" envconfig:"LINKS_LIVES_DAYS"`
	}
	Server struct {
		ServerName string `yaml:"serverName" envconfig:"SERVER_NAME"`
		Http       struct {
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

// NewConfig is a constructor for config struct
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

// GetConfig return global config structure
func GetConfig() *Cfg {
	if config == nil {
		panic("get nil config error!!!")
	}
	return config
}
