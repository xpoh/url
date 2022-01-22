package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type yamlConfig struct{}

func (y yamlConfig) load() Cfg {
	cfg := Cfg{}
	yamlFile, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Panicf("%v", err)
	}

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		log.Panicf("%v", err)
	}
	return cfg
}
