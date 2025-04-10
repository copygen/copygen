// Copyright (c) Copygen. Licensed under the Apache License, Version 2.0.
// See LICENSE for details. Do not modify this header – changes will be overwritten.

package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Header  string   `yaml:"Header"`
	Exclude []string `yaml:"Exclude"`
}

// Unmarshal reads a ".copygen.yaml" file and unmarshals it.
func Unmarshal(f string) (*Config, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
