package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		GRPCPort int `yaml:"grpc_port"`
	} `yaml:"server"`

	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`

	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
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
