package config

import (
	"go.uber.org/config"
	"go.uber.org/fx"
	"os"
)

const configPath = "./config/config.yml"

type AppCfg struct {
	fx.Out

	Provider config.Provider
}

func newConfig() (AppCfg, error) {
	cfg, err := os.Open(configPath)
	if err != nil {
		return AppCfg{}, err
	}

	provider, err := config.NewYAML(config.Source(cfg), config.Name("app_config"))
	if err != nil {
		return AppCfg{}, err
	}

	return AppCfg{Provider: provider}, nil
}
