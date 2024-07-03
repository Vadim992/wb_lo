package nats_streaming

import (
	"go.uber.org/config"
)

type Config struct {
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	WbChan    string `yaml:"wb_chan"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config

	err := provider.Get("nats_streaming").Populate(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
