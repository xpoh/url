// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config implement extern configuration
// for butty service
package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type yamlConfig struct{}

// load read config from yaml file
func (y yamlConfig) load() (Cfg, error) {
	cfg := Cfg{}
	yamlFile, err := os.ReadFile("config.yaml")

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
