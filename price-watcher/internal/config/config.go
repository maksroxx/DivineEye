package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Kafka struct {
		Brokers []string `yaml:"brokers"`
		Topic   []string `yaml:"topic"`
	} `yaml:"kafka"`

	Binance struct {
		Symbols []string `yaml:"symbols"`
	} `yaml:"binance"`
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
