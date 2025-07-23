package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"database"`

	Kafka struct {
		Brokers     []string `yaml:"brokers"`
		AlertTopic  string   `yaml:"alerts_topic"`
		PricesTopic string   `yaml:"prices_topic"`
	} `yaml:"kafka"`

	Fcm struct {
		ServerKey string `yaml:"server_key"`
	} `yaml:"fcm"`
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
