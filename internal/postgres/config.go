package postgres

import "go.uber.org/config"

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"db_name"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config

	err := provider.Get("postgres").Populate(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
