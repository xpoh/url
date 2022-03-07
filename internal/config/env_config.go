// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package config implement extern configuration
// for butty service
package config

import (
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct{}

// load read main config from ENV
func (e envConfig) load() (Cfg, error) {
	cfg := Cfg{}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
