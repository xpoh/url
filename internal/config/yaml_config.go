package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type yamlConfig struct{}

func (y yamlConfig) load() Cfg {
	cfg := Cfg{}
	yamlFile, err := os.ReadFile("./configs/config.yaml")

	if err != nil {
		log.Panicln(err)
	}

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		log.Panicln(err)
	}
	return cfg
}
