package server

import "go.uber.org/config"

type Config struct {
	Port string `yaml:"port"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config

	err := provider.Get("server").Populate(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
