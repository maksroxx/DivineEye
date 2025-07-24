package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		GRPCPort int `yaml:"grpc_port"`
	} `yaml:"server"`

	Database struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"database"`

	Kafka struct {
		Group   string   `yaml:"group"`
		Brokers []string `yaml:"brokers"`
		Topic   []string `yaml:"topics"`
	} `yaml:"kafka"`
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
