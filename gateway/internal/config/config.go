package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Address int `yaml:"port"`
	} `yaml:"server"`

	Services struct {
		Auth         string `yaml:"auth"`
		Alert        string `yaml:"alert"`
		Notification string `yaml:"notification"`
	} `yaml:"services"`
}

func Load(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
